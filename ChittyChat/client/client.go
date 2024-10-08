package main

import (
	proto "ChittyChat/grpc"
	"context"
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

	wg.Wait() // Wait for both clients to finish

}

func client(clientNumber int, wg *sync.WaitGroup) {
	for {
		defer wg.Done()

		conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("could not connect. plz help")
		}
		defer conn.Close()

		client := proto.NewChittychatDBClient(conn)

		/*post ,err := client.PublishPost(context.Background(), &proto.Post{})
		if err != nil {
			log.Fatalf("could not post")
		}*/
		post := &proto.Post{
			Post:        "I am so cool" + string(clientNumber),
			LamportTime: 32121321,
		}

		worked, err := client.PublishPost(context.Background(), post)
		if err != nil || worked.Posted == false {
			log.Fatalf("could not post")
		}

		posts, err := client.GetPosts(context.Background(), &proto.Empty{})
		if err != nil {
			log.Fatalf("could not get posts")
		}

		for _, post := range posts.Posts {
			println(" - "+post, clientNumber)
		}
		time.Sleep(time.Second * 10)
	}

}
