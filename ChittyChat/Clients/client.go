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

	chats := []string{
		"Excited for the weekend! 😎 #TGIF",
		"Just finished a great book! Highly recommend it. 📚",
		"Coffee first, adulting second. ☕️",
		"The sunset today was absolutely stunning! 🌅",
		"Anyone else obsessed with this new song? 🎧 #NowPlaying",
		"Feeling grateful for all the little things in life. 🙏",
		"Can’t wait to travel again! ✈️ #Wanderlust",
		"Working from home in pajamas is the best! 🧸 #RemoteLife",
		"Looking for Netflix recommendations. What should I watch next? 🎬",
		"New day, new goals. Let’s get it! 💪",
		"I just love autumn. The colors are so beautiful! 🍁🍂",
		"Is it just me, or did today feel extra long? ⏳",
		"I miss spontaneous road trips! 🚗",
		"Trying out a new recipe today. Fingers crossed! 🍳",
		"Can’t believe it’s almost the end of the year! 🎉",
		"Just ran my first 5K! Feeling amazing! 🏃‍♂️",
		"Let’s make kindness go viral. Be good to each other! 💛",
		"Monday motivation: Keep pushing forward! 🚀",
		"Grabbing brunch with friends this weekend. Can’t wait! 🥑🥂",
		"Loving this new podcast. So insightful! 🎙️ #PodcastRecommendation",
	}

	var wg sync.WaitGroup
	wg.Add(2) // We are launching two clients

	go client(1, &wg, chats)
	go client(2, &wg, chats)
	go client(3, &wg, chats)

	wg.Wait() // Wait for all clients to finish
}

func client(clientNumber int, wg *sync.WaitGroup, chats []string) {
	LamportTime := 1
	for {
		defer wg.Done()

		conn, err := grpc.NewClient("localhost:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatalf("Client", clientNumber, "could not connect")
		}
		defer conn.Close()

		client := proto.NewChittychatDBClient(conn)

		if LamportTime == 1 {
			client.Connect(context.Background(), &proto.ClientInfo{Cn: int64(clientNumber), LamportTime: int64(LamportTime)})
		}

		randomChatNr := rand.Intn(20)

		post := &proto.Post{ //Making a Post
			Post:        fmt.Sprintf(chats[randomChatNr]+", sent by clientNr: %d", clientNumber),
			LamportTime: int64(LamportTime),
		}

		// Sends the post to the server
		serverReturn, err := client.PublishPost(context.Background(), post)
		if err != nil || serverReturn.Posted == false {
			log.Fatalf("Client ", clientNumber, "Could not post")
		}
		LamportTime = int(serverReturn.LamportTime) //lamport time updates from what it recieved

		

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
