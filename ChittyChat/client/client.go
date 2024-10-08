package main

import (
	proto "ChittyChat/grpc"
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2) // We are launching two clients

	go client(1, &wg)
	go client(2, &wg)

	wg.Wait() // Wait for both clients to finish

}

func client(clientNumber int, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect. plz help")
	}
	defer conn.Close()

	client := proto.NewChittychatDBClient(conn)

	posts, err := client.GetPosts(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatalf("could not get posts")
	}
	for _, post := range posts.Posts {
		println(" - "+post, clientNumber)
	}
}
