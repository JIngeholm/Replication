package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"auction/proto/proto"
	auction "auction/proto/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuctionServer struct {
	auction.UnimplementedAuctionServiceServer
	mu             sync.Mutex
	NodeID         int64
	isPrimary      bool
	highestBid     int
	highestBidder  int
	auctionEndTime time.Time

	backupNodes []int64 // Replication: backup nodes
	bidders     []auction.AuctionServiceClient
}

func NewAuctionServer(id int64) *AuctionServer {
	return &AuctionServer{
		highestBid:     0,
		highestBidder:  0,
		auctionEndTime: time.Now().Add(20 * time.Second), // Auction ends in 100 seconds
		NodeID:         id,
	}
}

func (s *AuctionServer) Bid(ctx context.Context, req *auction.BidRequest) (*auction.BidResponse, error) {
	fmt.Println("bid 0")

	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println("bid 1")

	if s.isPrimary {
		s.RegisterBidder(ctx, &auction.BidderRequest{Address: fmt.Sprintf("localhost:%d", req.Bidder)})
	}
	fmt.Println("bid 2")

	// Auction is over
	if time.Now().After(s.auctionEndTime) {
		s.syncState()
		return &auction.BidResponse{Status: "fail"}, nil
	}
	fmt.Println("bid 3")

	// Check if the bid is higher than the previous highest bid
	if req.Amount <= int32(s.highestBid) {
		s.syncState()
		return &auction.BidResponse{Status: "fail"}, nil
	}
	fmt.Println("bid 4")

	// Register the new bid
	s.highestBid = int(req.Amount)
	s.highestBidder = int(req.Bidder)

	// Synchronize state with backup nodes
	s.syncState()
	fmt.Println("bid 5")

	return &auction.BidResponse{Status: "success"}, nil
}

func (s *AuctionServer) Result(ctx context.Context, req *auction.ResultRequest) (*auction.ResultResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Now().After(s.auctionEndTime) {
		return &auction.ResultResponse{
			Result: fmt.Sprintf("Winner: %d with bid %d", s.highestBidder, s.highestBid),
		}, nil
	}

	return &auction.ResultResponse{
		Result: fmt.Sprintf("Highest bid: %d by %d", s.highestBid, s.highestBidder),
	}, nil
}

// Synchronize auction state across backup nodes
func (s *AuctionServer) syncState() {
	fmt.Printf("Bidders: %d", len(s.bidders))
	fmt.Println()

	for _, bidder := range s.bidders {

		

		// Sending state to backup nodes (for simplicity, no response handling)
		bidder.Bid(context.Background(), &auction.BidRequest{
			Bidder: int32(s.highestBidder),
			Amount: int32(s.highestBid),
		})
	}
	fmt.Printf("highest bidder: %d", s.highestBidder)
		fmt.Println()

		fmt.Printf("highest bid: %d", s.highestBid)
		fmt.Println()
}

func (n *AuctionServer) getClient(server string) (proto.AuctionServiceClient, error) {
	conn, err := grpc.NewClient(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return proto.NewAuctionServiceClient(conn), nil
}

func (s *AuctionServer) RegisterBidder(ctx context.Context, req *auction.BidderRequest) (*auction.BidderResponse, error) {

	fmt.Println("reg 1")

	if !s.isPrimary {
		return &auction.BidderResponse{Status: "Fail: not primary"}, nil
	}
	fmt.Println("reg 2")

	bidder, err := s.getClient(req.Address)
	if err != nil {
		log.Fatalf("Error")
	}
	fmt.Println("reg 3")

	s.bidders = append(s.bidders, bidder)

	log.Printf("Bidder registered: %s", req.Address)
	return &auction.BidderResponse{Status: "succes"}, nil
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
	fmt.Println("test 1")

	port := flag.Int64("port", 50051, "Port for the gRPC server to listen on")
	flag.Parse()
	server := NewAuctionServer(*port)

	if *port == int64(50051) {
		fmt.Println("test 50051")

		server.isPrimary = true
		backup := NewAuctionServer(*port + 1)
		server.backupNodes = append(server.backupNodes, backup.NodeID)

	}
	fmt.Println("test 2")

	// Start the auction server (Primary node)
	go server.Start(fmt.Sprintf(":%d", server.NodeID))

	if !server.isPrimary {
		fmt.Println("test not primary")

		primary, err := server.getClient("localhost:50051")
		if err != nil {
			log.Fatalf("Error")
		}

		req := &auction.BidRequest{
			Bidder: int32(*port),
			Amount: 500,
		}

		if *port == 50053 {
			req.Amount = 600
		} else if *port == 50054 {
			req.Amount = 550
		}

		primary.Bid(context.Background(), req)

		time.Sleep(3 * time.Second)
		resReq := &auction.ResultRequest{}
		resp, err := (server.Result(context.Background(), resReq))
		if err != nil {
			log.Fatalf("Error")
		}
		log.Printf(resp.Result)
		fmt.Println("test 2.875")

	}
	fmt.Println("test 3")

	if server.isPrimary {
		for {
			time.Sleep(1 * time.Second)
			if time.Now().After(server.auctionEndTime) {
				resReq := &auction.ResultRequest{}
				resp, err := (server.Result(context.Background(), resReq))
				if err != nil {
					log.Fatalf("Error")
				}
				log.Printf(resp.Result)
				break
			}

		}
	}
	select {}
}
