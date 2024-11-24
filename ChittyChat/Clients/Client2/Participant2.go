package main

import (
	proto "ChittyChat/grpc"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	connectionLogLength int
	LamportTime         int
	clientNumber        int
	stopChan            chan struct{}
)

func main() {
	clientNumber = 2
	stopChan = make(chan struct{}) // Initialize once
	client()
	select {}
}

func client() {
	reader := bufio.NewReader(os.Stdin)
	LamportTime = 1
	isConnected := false
	connectionLogLength = 0

	conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client %d could not connect: %v", clientNumber, err)
	}
	defer conn.Close()

	client := proto.NewChittychatDBClient(conn)

	// Start listening to updates
	go listenToUpdates(client, stopChan)

	for {
		if !isConnected {
			// Signal the goroutine to stop
			stopChan <- struct{}{}

			isConnected = tryToConnectToServer(client, clientNumber, LamportTime, reader)

			// Restart goroutine after connecting
			go listenToUpdates(client, stopChan)
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		length := len(input)

		if input == "Disconnect" {
			client.Disconnect(context.Background(), &proto.ClientInfo{Cn: int64(clientNumber), LamportTime: int64(LamportTime)})
			isConnected = false
		} else if length > 128 {
			fmt.Println("Message too long!")
		} else {
			post := &proto.Post{
				Post:        input,
				LamportTime: int64(LamportTime),
			}

			serverReturn, err := client.PublishPost(context.Background(), post)
			if err != nil || !serverReturn.Posted {
				log.Fatalf("Client %d could not post: %v", clientNumber, err)
			}
			LamportTime = int(serverReturn.LamportTime)
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

func listenToUpdates(client proto.ChittychatDBClient, stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			return
		default:
			time.Sleep(time.Second)
			serverConnectionLog, err := client.GetConnectionLog(context.Background(), &proto.ClientInfo{Cn: int64(clientNumber), LamportTime: int64(LamportTime)})

			if err != nil {
				fmt.Println("Error fetching connection log:", err)
				continue
			}

			if len(serverConnectionLog.Logs) > connectionLogLength {
				for i := len(serverConnectionLog.Logs) - connectionLogLength; i > 0; i-- {
					fmt.Println(serverConnectionLog.Logs[len(serverConnectionLog.Logs)-i])
				}
				connectionLogLength = len(serverConnectionLog.Logs)
			}
		}
	}
}
