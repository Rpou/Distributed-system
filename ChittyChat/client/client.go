package main

import (
	proto "ChittyChat/grpc"
	"context"
	"fmt"
	"log"
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

	wg.Wait() // Wait for both clients to finish
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

		/*post ,err := client.PublishPost(context.Background(), &proto.Post{})
		if err != nil {
			log.Fatalf("could not post")
		}*/
		post := &proto.Post{
			Post:        fmt.Sprintf("I am so cool, sent by: %d", clientNumber),
			LamportTime: int64(LamportTime),
		}

		serverReturn, err := client.PublishPost(context.Background(), post)
		if err != nil || serverReturn.Posted == false {
			log.Fatalf("Client", clientNumber, "Could not post")
		}
		LamportTime = int(serverReturn.LamportTime) //lamport time updates from what it recieved

		posts, err := client.GetPosts(context.Background(), &proto.Empty{})
		if err != nil {
			log.Fatalf("Client", clientNumber, "Could not get posts")
		}

		println("Messages recieved by client:", clientNumber)
		for _, post := range posts.Posts {
			println(" - " + post)
		}
		time.Sleep(time.Second * 2)
	}

}
