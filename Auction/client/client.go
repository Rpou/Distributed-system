package main

import (
	proto "ITUServer/grpc"
	"context"
	"fmt"
	"time"

	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	randomBidNr := rand.Intn(100)
	go client(randomBidNr)

	randomBidNr = rand.Intn(100)
	go client(randomBidNr)

	select {}

}

func client(myBid int) {

	for {
		//connects to random client:
		randomNodeNr := rand.Intn(3)
		client := connectToNode(randomNodeNr)

		// see auction status
		auctionStatus, errorr := client.AuctionStatus(context.Background(), &proto.Empty{})
		if errorr != nil {
			fmt.Printf("Error fetching auction status from node %d: %v\n", randomNodeNr, errorr)
			time.Sleep(time.Millisecond * 100) // Retry with a different node
			continue
		}

		if auctionStatus.InProgress {
			// sends own bid, and recives auctions highest bid:
			highestbid, err := client.ClientRequest(context.Background(), &proto.ClientToNodeBid{
				Bid: int64(myBid),
			})
			if err != nil {
				fmt.Printf("Error sending bid to node %d: %v\n", randomNodeNr, err)
				time.Sleep(time.Millisecond * 100) // Retry with a different node
				continue
			}

			if highestbid.AuctionBid == int64(myBid) && highestbid.Giveacces {

				fmt.Println("I currently have the highest bid at:", myBid)

			} else {
				fmt.Println("I got rejected with:", myBid)
				randomAddedCash := rand.Intn(100) + 1
				myBid = myBid + randomAddedCash
			}

		} else {
			fmt.Println(auctionStatus.HighestBid)
			break
		}

		time.Sleep(time.Millisecond * 100)

	}
}

func connectToNode(nodeNumber int) proto.CommuncationClient {

	Node1FullAdd := "localhost:5051"
	Node2FullAdd := "localhost:5052"
	Node3FullAdd := "localhost:5053"

	var node proto.CommuncationClient
	var err error
	var conn *grpc.ClientConn

	for {

		if nodeNumber == 0 {

			conn, err = grpc.NewClient(Node1FullAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
			node = proto.NewCommuncationClient(conn)

		} else if nodeNumber == 1 {

			conn, err = grpc.NewClient(Node2FullAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
			node = proto.NewCommuncationClient(conn)

		} else {
			conn, err = grpc.NewClient(Node3FullAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
			node = proto.NewCommuncationClient(conn)

		}
		//if no errors, return node. If there is error it tries again
		if err == nil {
			fmt.Printf("Connected to node %d\n", nodeNumber)
			if conn != nil {
				return node
			}
		}

		fmt.Printf("Failed to connect to node %d: %v\n", nodeNumber, err)
		nodeNumber = rand.Intn(3) // Retry with a different node
	}
}
