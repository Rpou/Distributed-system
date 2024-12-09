package main

import (
	proto "ITUServer/grpc"
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	client()
	select {}

}

func client() {
	MyBid := 0
	reader := bufio.NewReader(os.Stdin)
	highbid := -1
	fmt.Println("Write a bid to participate in the auction. Write '{bid amount}' to bid or 'Status' to see auction status.")
	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		//connects to random client:
		randomNodeNr := rand.Intn(3)
		client, nodeNumber := connectToNode(randomNodeNr)

		fmt.Println("You connected to Node:", nodeNumber)
		if _, err := strconv.Atoi(input); err == nil {
			bid, err := strconv.Atoi(input)

			if err != nil {
				fmt.Printf("Error sending bid to node %d: %v\n", randomNodeNr, err)
				continue
			}

			if MyBid > bid {
				fmt.Println("this bid is invalid, as it is lower than your current bid, or was lower than 0")
			} else {
				MyBid = bid
				Output, err := client.Bid(context.Background(), &proto.ClientToNodeBid{
					Bid: int64(bid),
				})

				if err != nil {
					fmt.Printf("Error sending bid to node %d: %v\n", randomNodeNr, err)
					continue
				}

				fmt.Println(Output)
			}

		} else if input == "Status" {

			auctionStatus, errorr := client.Result(context.Background(), &proto.Empty{})
			if errorr != nil {
				fmt.Printf("Error fetching auction status from node %d: %v\n", randomNodeNr, errorr)
				time.Sleep(time.Millisecond * 100) // Retry with a different node
				continue
			}
			highbid = int(auctionStatus.HighestBid)
			if auctionStatus.InProgress {
				if MyBid == int(highbid) {
					fmt.Println("You have the highest bid!")
				} else {
					fmt.Println("The Auction is still ongoing! The highest bid is:", highbid)
				}

			} else {
				if MyBid == int(highbid) {
					fmt.Println("The Auction is over! You had the highest bid! You won!")
				} else {
					fmt.Println("The Auction is over! The highest bid was:", highbid)
				}
			}

		} else {
			fmt.Println("I dont understand the input. Please use the instruction at the top")
		}

	}
}

func connectToNode(nodeNumber int) (proto.CommuncationClient, int) {

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
			fmt.Printf("Connected to node %d\n", nodeNumber+1)
			if conn != nil {
				return node, nodeNumber + 1
			}
		}

		fmt.Printf("Failed to connect to node %d: %v\n", nodeNumber)
		nodeNumber = rand.Intn(3) // Retry with a different node
	}
}
