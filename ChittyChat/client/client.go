package main

import (
	proto "ChittyChat/grpc"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2) // We are launching two clients

	go client(1, &wg)
	go client(2, &wg)
	go client(3, &wg)

	wg.Wait() // Wait for all clients to finish
}

func client(clientNumber int, wg *sync.WaitGroup) {
	LamportTime := 1
	for {
		defer wg.Done()

		conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Client", clientNumber, "could not connect")
		}
		defer conn.Close()

		client := proto.NewChittychatDBClient(conn)

		post := &proto.Post{ //Making a Post
			Post:        fmt.Sprintf("I am so cool, sent by clientNr: %d", clientNumber),
			LamportTime: int64(LamportTime),
		}

		// Sends the post to the server
		serverReturn, err := client.PublishPost(context.Background(), post)
		if err != nil || serverReturn.Posted == false {
			log.Fatalf("Client", clientNumber, "Could not post")
		}
		LamportTime = int(serverReturn.LamportTime) //lamport time updates from what it recieved

		// retrives all the posts from the server
		posts, err := client.GetPosts(context.Background(), &proto.ClientLT{LamportTime: int64(LamportTime)})
		if err != nil {
			log.Fatalf("Client", clientNumber, "Could not get posts")
		}

		println("Messages recieved by client:", clientNumber)
		for _, post := range posts.Posts {
			println(" - " + post)
		}

		//20% chance it disconnects everytime it has made a post.
		randomNumber := rand.Intn(5) + 1
		if randomNumber == 1 {
			client.Disconnect(context.Background(), &proto.ClientInfo{Cn: int64(clientNumber), LamportTime: int64(LamportTime)})
			for {
				// Sleeps for 1 sec, and then has a 20% chance of connecting again. if it fails, it will try again after 1 sec
				time.Sleep(time.Second)
				randomNumber = rand.Intn(5) + 1
				if randomNumber == 2 {
					client.Connect(context.Background(), &proto.ClientInfo{Cn: int64(clientNumber), LamportTime: int64(LamportTime)})
					break
				}
			}
		}

		time.Sleep(time.Second * 2)
	}

}
