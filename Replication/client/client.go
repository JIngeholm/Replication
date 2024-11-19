package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	auction "auction/proto/proto"
	
)

func main() {
	// Connect to the auction server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}
	defer conn.Close()

	client := auction.NewAuctionServiceClient(conn)

	// Place a bid
	bid := &auction.BidRequest{Bidder: "Alice", Amount: 100}
	_, err = client.Bid(context.Background(), bid)
	if err != nil {
		log.Fatalf("could not bid: %v", err)
	}

	// Place another bid
	bid = &auction.BidRequest{Bidder: "Bob", Amount: 150}
	_, err = client.Bid(context.Background(), bid)
	if err != nil {
		log.Fatalf("could not bid: %v", err)
	}

	// Query the result
	resp, err := client.Result(context.Background(), &auction.ResultRequest{})
	if err != nil {
		log.Fatalf("could not get result: %v", err)
	}

	fmt.Println("Auction Result:", resp.Result)
}
