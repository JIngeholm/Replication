package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os" // Importing the os package
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
	isBackup       bool
	highestBid     int
	highestBidder  int
	auctionEndTime time.Time
	primary        auction.AuctionServiceClient
	backupClient   auction.AuctionServiceClient

	bidders     []auction.AuctionServiceClient
	biddersMap  map[string]bool // A map to check for duplicates
}


func NewAuctionServer(id int64) *AuctionServer {
	return &AuctionServer{
		highestBid:     0,
		highestBidder:  0,
		auctionEndTime: time.Now().Add(30 * time.Second),
		NodeID:         id,
		biddersMap:     make(map[string]bool), // Initialize the map here
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
        // Skip sending bid to self
        if bidder == s.primary {
            continue
        }

        _, err := bidder.Bid(context.Background(), &auction.BidRequest{
            Bidder: int32(s.highestBidder),
            Amount: int32(s.highestBid),
        })
        if err != nil {
            log.Printf("Error syncing state with bidder: %v", err)
        }
    }

	fmt.Println()
    log.Printf("highest bidder: %d", s.highestBidder)
    log.Printf("highest bid: %d", s.highestBid)
}


func (n *AuctionServer) getClient(server string) (proto.AuctionServiceClient, error) {
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewAuctionServiceClient(conn), nil
}

func (s *AuctionServer) RegisterBidder(ctx context.Context, req *auction.BidderRequest) (*auction.BidderResponse, error) {
	if !s.isPrimary {
		return &auction.BidderResponse{Status: "Fail: not primary"}, nil
	}

	// Use a map to check for duplicates (using Address as the key)
	bidderAddress := req.Address
	if s.biddersMap[bidderAddress] {
		return &auction.BidderResponse{Status: "Fail: bidder already registered"}, nil
	}

	// If not already in the map, register the bidder
	bidder, err := s.getClient(bidderAddress)
	if err != nil {
		log.Fatalf("Error registering bidder: %v", err)
	}

	// Add the bidder to the list and mark as registered in the map
	s.bidders = append(s.bidders, bidder)
	s.biddersMap[bidderAddress] = true

	fmt.Println()
	log.Printf("Bidder registered: %s", bidderAddress)
	return &auction.BidderResponse{Status: "success"}, nil
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

func (s *AuctionServer) SendHeartbeat(ctx context.Context, req *auction.HeartbeatRequest) (*auction.HeartbeatResponse, error) {
	return &auction.HeartbeatResponse{Status: "alive"}, nil
}

func (s *AuctionServer) CheckHeartbeat() {
	if s.isPrimary {
		// Primary doesn't need to check its own heartbeat
		return
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if s.primary == nil {
			log.Fatal("Primary is nil, unable to send heartbeat")
		}
		_, err := s.primary.SendHeartbeat(context.Background(), &auction.HeartbeatRequest{})
		if err != nil {
			s.mu.Lock()
			fmt.Println()
			if(s.NodeID == 50052){
				log.Println("Primary server '50051' not responding. Taking over as primary.")
			}else{
				log.Println("Primary server '50051' not responding. Promoting backup '50052' to primary.")
			}

			if(s.NodeID == 50052){
				s.isPrimary = true
				s.isBackup = false
				s.mu.Unlock()
				return
			}
		
			time.Sleep(3 * time.Second) // wait for backup to promote itself
			s.primary = s.backupClient
			req := &auction.BidRequest{
				Bidder: int32(s.highestBidder),
				Amount:int32(s.highestBid),
			}
			s.mu.Unlock()

			fmt.Println()
			log.Println("Resending highest bid and bidder to new primary")
			s.primary.Bid(context.Background(), req)
		
			break
		}
	}
}

func main() {
	port := flag.Int64("port", 50051, "Port for the gRPC server to listen on")
	flag.Parse()
	server := NewAuctionServer(*port)

	if *port == int64(50051) { // primary
		go server.Start(":50051")
		server.isPrimary = true
	}else if *port == int64(50052){ // backup
		go server.Start(":50052")
		go server.CheckHeartbeat()
		server.isBackup = true
		primary, err := server.getClient("localhost:50051")
		if err != nil {
			log.Fatalf("Error connecting to primary: %v", err)
		}
		server.primary = primary
	}

	if !server.isPrimary && !server.isBackup { // Script for bidding
		go server.Start(fmt.Sprintf(":%d", server.NodeID))
		primary, err := server.getClient("localhost:50051")
		if err != nil {
			log.Fatalf("Error connecting to primary: %v", err)
		}
		backupClient, err := server.getClient("localhost:50052")
		if err != nil {
			log.Fatalf("Error connecting to backup: %v", err)
		}
		
		server.primary = primary
		server.backupClient = backupClient

		go server.CheckHeartbeat()

		//First bid
		req := &auction.BidRequest{
			Bidder: int32(*port),
			Amount: 500,
		}

		if *port == 50053 {
			req.Amount = 600
		} else if *port == 50054 {
			req.Amount = 550
		}

		server.primary.Bid(context.Background(), req)

		time.Sleep(3 * time.Second)

		//Result from bid
		resReq := &auction.ResultRequest{}
		resp, err := (server.Result(context.Background(), resReq))
		if err != nil {
			log.Fatalf("Error getting result: %v", err)
		}
		fmt.Println()
		log.Printf(resp.Result)
	}

	if(server.isPrimary){
		// Crash the primary server manually
		time.Sleep(10 * time.Second)
		os.Exit(1) 
	}

	if !server.isPrimary && !server.isBackup {
		time.Sleep(10 * time.Second)

		//Query auction before it is over
		resReq := &auction.ResultRequest{}
		resp, err := (server.Result(context.Background(), resReq))
		if err != nil {
			log.Fatalf("Error getting result: %v", err)
		}
		fmt.Println()
		fmt.Println("Querying the auction:")
		log.Printf("%d %s", *port, resp.Result)

		//Bidding higher than first bid
		if *port != 50053 && *port != 50054 {
			req := &auction.BidRequest{
				Bidder: int32(*port),
				Amount: 650,
			}
			fmt.Println("Making bigger bid")
			server.primary.Bid(context.Background(), req)
		}

		time.Sleep(35 * time.Second)

		//Query auction after it is over
		resp, err = (server.Result(context.Background(), resReq))
		if err != nil {
			log.Fatalf("Error getting result: %v", err)
		}
		fmt.Println()
		fmt.Println("Querying the auction:")
		log.Printf("%d %s", *port, resp.Result)
	}

	if server.isPrimary { //Primary gets result after auction is over
		for {
			time.Sleep(1 * time.Second)
			if time.Now().After(server.auctionEndTime) {
				resReq := &auction.ResultRequest{}
				resp, err := (server.Result(context.Background(), resReq))
				if err != nil {
					log.Fatalf("Error getting result: %v", err)
				}
				log.Printf(resp.Result)
				break
			}
		}
	}

	select {}
}
