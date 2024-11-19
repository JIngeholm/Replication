package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	auction "auction/proto/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

type AuctionServer struct {
	auction.UnimplementedAuctionServiceServer
	mu             sync.Mutex
	isPrimary      bool
	highestBid     int
	highestBidder  string
	auctionEndTime time.Time
	backupNodes    []auction.AuctionServiceClient // Replication: backup nodes
}

func NewAuctionServer() *AuctionServer {
	return &AuctionServer{
		highestBid:     0,
		highestBidder:  "",
		auctionEndTime: time.Now().Add(100 * time.Second), // Auction ends in 100 seconds
	}
}

func (s *AuctionServer) Bid(ctx context.Context, req *auction.BidRequest) (*auction.BidResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Auction is over
	if time.Now().After(s.auctionEndTime) {
		return &auction.BidResponse{Status: "fail"}, nil
	}

	// Check if the bid is higher than the previous highest bid
	if req.Amount <= int32(s.highestBid) {
		return &auction.BidResponse{Status: "fail"}, nil
	}

	// Register the new bid
	s.highestBid = int(req.Amount)
	s.highestBidder = req.Bidder

	// Synchronize state with backup nodes
	s.syncState()

	return &auction.BidResponse{Status: "success"}, nil
}

func (s *AuctionServer) Result(ctx context.Context, req *auction.ResultRequest) (*auction.ResultResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Now().After(s.auctionEndTime) {
		return &auction.ResultResponse{
			Result: fmt.Sprintf("Winner: %s with bid %d", s.highestBidder, s.highestBid),
		}, nil
	}

	return &auction.ResultResponse{
		Result: fmt.Sprintf("Highest bid: %d by %s", s.highestBid, s.highestBidder),
	}, nil
}

// Synchronize auction state across backup nodes
func (s *AuctionServer) syncState() {
	for _, backup := range s.backupNodes {
		// Sending state to backup nodes (for simplicity, no response handling)
		backup.Bid(context.Background(), &auction.BidRequest{
			Bidder: s.highestBidder,
			Amount: int32(s.highestBid),
		})
	}
}

// Start the gRPC server
func (s *AuctionServer) Start(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	auction.RegisterAuctionServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// Create the primary auction node
	server := NewAuctionServer()

	// Start the auction server (Primary node)
	go server.Start(":50051")

	// Simulate backup nodes by creating AuctionServiceClients (in real implementation, these would run on different machines)
	// Here, we connect to the primary server for simplicity
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := auction.NewAuctionServiceClient(conn)
	server.backupNodes = append(server.backupNodes, client)

	// Simulate auction clients sending bids and querying results
	time.Sleep(1 * time.Second) // Allow server to start

	// Simulating bids
	_, err = client.Bid(context.Background(), &auction.BidRequest{Bidder: "Alice", Amount: 100})
	if err != nil {
		log.Fatalf("could not bid: %v", err)
	}

	_, err = client.Bid(context.Background(), &auction.BidRequest{Bidder: "Bob", Amount: 150})
	if err != nil {
		log.Fatalf("could not bid: %v", err)
	}

	// Query the auction result
	resp, err := client.Result(context.Background(), &auction.ResultRequest{})
	if err != nil {
		log.Fatalf("could not get result: %v", err)
	}

	fmt.Println("Auction Result:", resp.Result)
}
