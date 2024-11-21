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

	backupNodes []AuctionServer
	bidders     []auction.AuctionServiceClient
}

func NewAuctionServer(id int64) *AuctionServer {
	return &AuctionServer{
		highestBid:     0,
		highestBidder:  0,
		auctionEndTime: time.Now().Add(20 * time.Second),
		NodeID:         id,
	}
}

func (s *AuctionServer) Bid(ctx context.Context, req *auction.BidRequest) (*auction.BidResponse, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isPrimary {
		s.RegisterBidder(ctx, &auction.BidderRequest{Address: fmt.Sprintf("localhost:%d", req.Bidder)})
	}

	if time.Now().After(s.auctionEndTime) {
		s.syncState()
		return &auction.BidResponse{Status: "fail"}, nil
	}

	if req.Amount <= int32(s.highestBid) {
		s.syncState()
		return &auction.BidResponse{Status: "fail"}, nil
	}

	s.highestBid = int(req.Amount)
	s.highestBidder = int(req.Bidder)

	s.syncState()

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

func (s *AuctionServer) syncState() {

	for _, bidder := range s.bidders {

		bidder.Bid(context.Background(), &auction.BidRequest{
			Bidder: int32(s.highestBidder),
			Amount: int32(s.highestBid),
		})
	}

	fmt.Println()
	log.Printf("highest bidder: %d", s.highestBidder)
	log.Printf("highest bid: %d", s.highestBid)
}

func (n *AuctionServer) getClient(server string) (proto.AuctionServiceClient, error) {
	fmt.Println(server)

	conn, err := grpc.NewClient(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewAuctionServiceClient(conn), nil
}

func (s *AuctionServer) RegisterBidder(ctx context.Context, req *auction.BidderRequest) (*auction.BidderResponse, error) {

	if !s.isPrimary {
		return &auction.BidderResponse{Status: "Fail: not primary"}, nil
	}

	bidder, err := s.getClient(req.Address)
	if err != nil {
		log.Fatalf("Error")
	}

	s.bidders = append(s.bidders, bidder)
	for _, backup := range s.backupNodes {
		backup.bidders = append(backup.bidders, bidder)
	}
	log.Printf("Bidder registered: %s", req.Address)
	return &auction.BidderResponse{Status: "succes"}, nil
}

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

	port := flag.Int64("port", 50051, "Port for the gRPC server to listen on")
	flag.Parse()
	server := NewAuctionServer(*port)

	if *port == int64(50051) {

		server.isPrimary = true
		backup := NewAuctionServer(50052)
		/*
			backupclient, err := backup.getClient("localhost:50052")
			if err != nil {
				log.Fatalf("Error")
			}
		*/
		server.backupNodes = append(server.backupNodes, *backup)

	}

	go server.Start(fmt.Sprintf(":%d", server.NodeID))

	if !server.isPrimary {

		primary, err := server.getClient("localhost:50051")
		if err != nil {
			log.Fatalf("Error")
		}
		/*
			backupClient, err := server.getClient("localhost:50052")
			if err != nil {
				log.Fatalf("Error")
			}
		*/
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

	}

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
