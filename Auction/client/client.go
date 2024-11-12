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

	randomNodeNr := rand.Intn(3)

	client := connectToNode(randomNodeNr)

	

	client.ClientRequest(context.Background(), &proto.ClientToNodeBid{
		Bid: 100,
	})

	fmt.Print(client)

	client.getbid

}

func connectToNode(nodeNumber int) proto.CommuncationClient {

	Node1FullAdd := "localhost:5051"
	Node2FullAdd := "localhost:5052"
	Node3FullAdd := "localhost:5053"

	var client proto.CommuncationClient
	var err error
	var conn *grpc.ClientConn

	if nodeNumber == 0 {

		conn, err = grpc.NewClient(Node1FullAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
		client = proto.NewCommuncationClient(conn)

	} else if nodeNumber == 1 {

		conn, err = grpc.NewClient(Node2FullAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
		client = proto.NewCommuncationClient(conn)

	} else {

		conn, err = grpc.NewClient(Node3FullAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
		client = proto.NewCommuncationClient(conn)
	}

	if err != nil {

	}

	return client
}
