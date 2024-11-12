package main

import (
	proto "ITUServer/grpc"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	CurrentHighestBid = 0
	TimeLeftOfAuction = 100;
)

type CommuncationServer struct {
	proto.UnimplementedCommuncationServer
	timestamp  int
	id         int
	mu         sync.Mutex
	wantAccess bool
}

func (s *CommuncationServer) Request(ctx context.Context, in *proto.CriticalData) (*proto.Accept, error) {

	// Accept if the other timestamp is smaller than this clients timestamp, or if this client does not want access

	return &proto.Accept{Giveacces: in.Time > int64(s.timestamp) || !s.wantAccess}, nil
}

func main() {
	server1 := &CommuncationServer{id: 1}
	server2 := &CommuncationServer{id: 2}
	server3 := &CommuncationServer{id: 3}

	Node1Add := ":5051"
	Node2Add := ":5052"
	Node3Add := ":5053"

	Node1FullAdd := "localhost:5051"
	Node2FullAdd := "localhost:5052"
	Node3FullAdd := "localhost:5053"

	go server1.start_server(Node1Add)
	go server2.start_server(Node2Add)
	go server3.start_server(Node3Add)

	go server1.auction(Node2FullAdd, Node3FullAdd)
	go server2.auction(Node1FullAdd, Node3FullAdd)
	go server3.auction(Node1FullAdd, Node2FullAdd)

	// Keep the main function alive to prevent exit
	select {}

}

func (s *CommuncationServer) start_server(NodeAddress string) {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", NodeAddress)
	if err != nil {
		log.Fatalf("Did not work")
	}

	proto.RegisterCommuncationServer(grpcServer, s)

	fmt.Println("opened Node")

	if err != nil {
		log.Fatalf("Did not work")
	}

	err = grpcServer.Serve(listener)
}

func (s *CommuncationServer) auction(peer1 string, peer2 string) {
	s.timestamp = s.id

	for {
	
		if true { //bid has been made
			for {
				conn, err := grpc.NewClient(peer1, grpc.WithTransportCredentials(insecure.NewCredentials()))
				conn2, err := grpc.NewClient(peer2, grpc.WithTransportCredentials(insecure.NewCredentials()))
				defer conn.Close()
				defer conn2.Close()
				if err != nil {
					log.Fatalf("Connection failed")
				}

				accept1 := getPeerConnection(conn, int64(s.timestamp))
				accept2 := getPeerConnection(conn2, int64(s.timestamp))
				if accept1 && accept2 {
					if  true { //bid is higher than current
						fmt.Println("I am node ", s.id, " Current new price: ", , " timestamp: ", s.timestamp)
						
					}  
					//send current highest back

					time.Sleep(time.Millisecond * 100)
					break
				} else {
					fmt.Println("I am ", s.id, " I got no access granted ", s.timestamp)
					time.Sleep(time.Millisecond * 100)
					s.timestamp += 5
				}
			}
		} else {
			time.Sleep(time.Millisecond * 100)
		}

		if TimeLeftOfAuction < 0 {
			break;
		}

	}

	//Returner endelig pris til client 
}

func getPeerConnection(conn *grpc.ClientConn, timestamp int64) bool {
	client := proto.NewCommuncationClient(conn)

	accept, err := client.Request(context.Background(), &proto.CriticalData{
		CriticalData: 2,
		Time:         timestamp,
	})

	if err != nil {
		log.Fatalf("Request failed")
	}

	return accept.Giveacces

}
