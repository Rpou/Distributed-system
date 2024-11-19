package main

import (
	proto "ITUServer/grpc"
	"context"
	"fmt"

	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	randomBidNr := rand.Intn(100)
	go client(randomBidNr)

	randomBidNr = rand.Intn(100)
	go client(randomBidNr)

}

func client(myBid int){

	for{
		//connects to random client:
		randomNodeNr := rand.Intn(3)
		client := connectToNode(randomNodeNr)

		client.Result(context.Background(), &proto.Empty{})

		// sends own bid, and recives auctions highest bid:
		highestbid, err := client.ClientRequest(context.Background(), &proto.ClientToNodeBid{
			Bid: int64(myBid),
		})
		if err != nil {

		}

		if highestbid.AuctionBid == int64(myBid){

			fmt.Println("I currently have the highest bid at:",myBid)

		} else {

			randomAddedCash := rand.Intn(100) + 1
			myBid = myBid + randomAddedCash
		}
	}
}

func connectToNode(nodeNumber int) proto.CommuncationClient {

	Node1FullAdd := "localhost:5051"
	Node2FullAdd := "localhost:5052"
	Node3FullAdd := "localhost:5053"

	var node proto.CommuncationClient
	var err error
	var conn *grpc.ClientConn

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

	if err != nil {

	}

	return node
}
