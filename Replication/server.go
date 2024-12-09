package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os" 
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




// Bid is called on the primary server by the other nodes using their grpc connection they got from the getClient method



func (s *AuctionServer) Bid(ctx context.Context, req *auction.BidRequest) (*auction.BidResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()




	// Registers new bidders, i.e. adds them to the primary server's list "bidders"

	if s.isPrimary {
		s.RegisterBidder(ctx, &auction.BidderRequest{Address: fmt.Sprintf("localhost:%d", req.Bidder)})
	}

	// Check if auction is over


	if time.Now().After(s.auctionEndTime) {
		s.syncState()
		return &auction.BidResponse{Status: "fail: auction over"}, nil
	}


	// Checks if bid is high enough


	if req.Amount <= int32(s.highestBid) {
		s.syncState()
		return &auction.BidResponse{Status: "fail: bid not high enough"}, nil
	}

	s.highestBid = int(req.Amount)
	s.highestBidder = int(req.Bidder)

	// Calls a method which propagates the highest "bidder" and "bid" to all the nodes in the list "bidders" called SyncState
	// Ensuring Replication, since all the bidders will have a replica of the auction state
	// This only propegates the auction state when a new bid is made.


	s.syncState()

	return &auction.BidResponse{Status: "success"}, nil
}




// This is the method which handles printing the result of a bid in all of the node's terminals

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



// Propagates the state of the auction to all node's in the list "bidders"
// This ensures all the nodes, have a replica of the "state of the auction" i.e. the highest "bidder" and "bid"


func (s *AuctionServer) syncState() {

    for _, bidder := range s.bidders {
        // Skip sending bid to self
        if bidder == s.primary {
            continue
        }


		// Calls the bid method on each of the bidders in the map
		// In the bid method each node then assigns the values "highest bidder" and "highest bid"
		// To be equal to the values that this method propagates


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




// Using grpc we establish a connection between nodes using this method getClient



func (n *AuctionServer) getClient(server string) (proto.AuctionServiceClient, error) {
	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewAuctionServiceClient(conn), nil
}




// This method RegisterBidder is used by the primary server and it adds nodes to the list "bidders" 
// This is the map which the method SyncState uses to propagate the state of the auction to the other nodes
// Ensuring replication :)


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

	// Add the bidder to the map and mark as registered in the map
	s.bidders = append(s.bidders, bidder)
	s.biddersMap[bidderAddress] = true

	fmt.Println()
	log.Printf("Bidder registered: %s", bidderAddress)
	return &auction.BidderResponse{Status: "success"}, nil
}



// Starts the server object




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



// These methods run as a thread in the background, and they checks on the primary server, using the connection between the nodes and the primary server
// They send a signal "heartbeat" to the primary server and if the primary doesn't respond, it promotes the backup server, which it is already connected to
// via the "backUpClient" connection, to the new primary. 
// So this handles the system failure of at least one node, the only node which would completely crash the entire system if it crashes would be the primary node, and we only
// need to account for one failure, so checking the primary is all we need to do. And if it fails we switch the backup node to be the primary.



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

		// Sends heartbeat signal to the primary
		_, err := s.primary.SendHeartbeat(context.Background(), &auction.HeartbeatRequest{})
		if err != nil {

			// Creates a critical section ensure that no one makes a bid once they have noticed the primary is down

			s.mu.Lock()
			fmt.Println()
			if(s.NodeID == 50052){
				log.Println("Primary server '50051' not responding. Taking over as primary.")
			}else{
				log.Println("Primary server '50051' not responding. Promoting backup '50052' to primary.")
			}

			// Promotes the backup node '50052' to be the the new primary server.
			// This happens by simply changing a boolean value 

			if(s.NodeID == 50052){
				s.isPrimary = true
				s.isBackup = false
				s.mu.Unlock()
				return
			}


			// If you are not the backup you wait for the time it takes for the backup to promote itself
			// When that time has passed you change the connnection "primary" which is a connection each node has
			// to the primary node, to be equal to the connection "backUpClient" which is the connection each node has,
			// to the back up node.

			// In other words the backup is promoted by every node to be the primary

			// It then sends the replica of the auction state to the primary node.
			// The values "highestBid" and "highestBidder" are replicas of the auction state ensuring the auction state can be
			// as it was before the crash. This is how we implement replication.

		
			time.Sleep(3 * time.Second) // wait for backup to promote itself
			s.primary = s.backupClient
			req := &auction.BidRequest{
				Bidder: int32(s.highestBidder),
				Amount:int32(s.highestBid),
			}
			s.mu.Unlock()


			// Resends the highest bid i.e. the replicate of the action state to the new primary


			fmt.Println()
			log.Println("Resending highest bid and bidder to new primary")
			s.primary.Bid(context.Background(), req)
		
			break
		}
	}
}



// The main function serves as a script essentially. Depending on what port the node started on, the node will have different behaviours.
// 50051 is the initial primary and therefore sets it's boolean value "isPrimary" to true.
// The backup, 50052, starts a thread "CheckHeartbeat" which we explained above, ensuring that if the primary fails it will take over.
// All the other nodes will run the same thread to esure the program can handle a crash of the primary server.
// Just like the other which i will describe below, it also establishes a connection to the primary using the getClient method,
// The connections is called "primary"




func main() {
	port := flag.Int64("port", 50051, "Port for the gRPC server to listen on")
	flag.Parse()
	server := NewAuctionServer(*port)

	if *port == int64(50051) { // primary
		go server.Start(":50051")

		// Sets itself to be primary by settings "isPrimary" to true

		server.isPrimary = true
	}else if *port == int64(50052){ // backup
		go server.Start(":50052")
		go server.CheckHeartbeat()

		// Sets itself to be the backup by setting "isBackup" to true

		server.isBackup = true

		// Establishes a connection to the primary called "primary"

		primary, err := server.getClient("localhost:50051")
		if err != nil {
			log.Fatalf("Error connecting to primary: %v", err)
		}
		server.primary = primary
	}



	// This is where the script begins for all nodes
	// Depending on who you are: primary,backup or another node, your behaviour is defined here
	// Just like the backup node above the other nodes establish a connection with the primary called "primary"
	// using the getClient method. They also establish a connection with the backup server called "backUpClient".
	// They will change the "primary" connection to be equal to the backup connection if the primary server fails




	// I have added this next part before line: 367, so that the primary node adds the backup node to it's list "bidders"
	// We also keep track of what nodes are in our list of bidders by using the "biddersMap" which made it easier to simply look up bool values
	
	if(server.isPrimary){
		server.bidders = append(server.bidders, server.backupClient)
		server.biddersMap["localhost:50052"] = true
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


		// Here all the other nodes start a thread checking on the primary server just as the backup server does above


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


    // Now the primary will manually crash itself and when this happens as i explained above, the other nodes will
	// promote the backup server to be the primary. They do this by changing their connection "primary" to be equal to
	// their connection "backUpClient"
	// I know I'm explaining things multiple times, but i want to make sure the message is coming through


	if(server.isPrimary){
		// Crash the primary server manually
		time.Sleep(10 * time.Second)
		os.Exit(1) 
	}

	// Rememeber the backup server's boolean "isPrimary" is now true and therefore acts as the primary once did

	// Here we make more bids and queries

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


	// The primary which in this case is now the former backup which became the primary, asks for the results of the auction and logs the statement


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
