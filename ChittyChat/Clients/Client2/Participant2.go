package main

import (
	proto "ChittyChat/grpc"
	"context"
	"fmt"
	"log"
	"strings"

	"bufio"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	connectionLog 	  []string
)

func main() {

	client(2)
	select {}
}

func client(clientNumber int) {
	reader := bufio.NewReader(os.Stdin)
	LamportTime := 1
	isConnected := false

	conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Client", clientNumber, "could not connect")
	}
	defer conn.Close()

	client := proto.NewChittychatDBClient(conn)
	for {

		if !isConnected {
			isConnected = tryToConnectToServer(client, clientNumber, LamportTime, reader)
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("i got an error")
		}

		input = strings.TrimSpace(input)
		length := len(input)

		if input == "Disconnect" {
			client.Disconnect(context.Background(), &proto.ClientInfo{Cn: int64(clientNumber), LamportTime: int64(LamportTime)})
			isConnected = false
		} else if length > 128 {
			fmt.Println("Too long message!")
		} else {

			post := &proto.Post{ //Making a Post
				Post:        fmt.Sprintf(input),
				LamportTime: int64(LamportTime),
			}

			// Sends the post to the server
			serverReturn, err := client.PublishPost(context.Background(), post)
			if err != nil || serverReturn.Posted == false {
				log.Fatalf("Client ", clientNumber, "Could not post")
			}
			LamportTime = int(serverReturn.LamportTime) //lamport time updates from what it recieved

		}
	}

}

func tryToConnectToServer(client proto.ChittychatDBClient, clientNumber int, LamportTime int, reader *bufio.Reader) bool {

	fmt.Println("Waiting for input... Type 'Join' to proceed.")
	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "Join" {
			fmt.Println("Join command accepted. Joining...")
			client.Connect(context.Background(), &proto.ClientInfo{Cn: int64(clientNumber), LamportTime: int64(LamportTime)})
			return true
		} else {
			fmt.Println("Invalid input. Type 'Join' to proceed.")
		}
	}
}
