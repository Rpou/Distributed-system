package main

import (
	proto "ITUServer/grpc"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	currentHighestBid = 0
	timeLeftOfAuction = 1000
)

type CommuncationServer struct {
	proto.UnimplementedCommuncationServer
	timestamp    int
	id           int
	otherServer1 string
	otherServer2 string
	mu           sync.Mutex
	wantAccess   bool
	status       string
}

func (s CommuncationServer) Request(ctx context.Context, in *proto.RequestAccess) (*proto.AcceptNodeRequest, error) {

	// Accept if the other timestamp is smaller than this clients timestamp, or if this client does not want access

	return &proto.AcceptNodeRequest{Giveacces: in.Time > int64(s.timestamp) || !s.wantAccess, Status: s.status}, nil
}

func (s *CommuncationServer) ClientRequest(ctx context.Context, in *proto.ClientToNodeBid) (*proto.AcceptClientRequest, error) {
	for {
		fmt.Println(s.id, "started asking for access")
		s.wantAccess = true
		conn, err := grpc.NewClient(s.otherServer1, grpc.WithTransportCredentials(insecure.NewCredentials()))
		conn2, err := grpc.NewClient(s.otherServer2, grpc.WithTransportCredentials(insecure.NewCredentials()))
		defer conn.Close()
		defer conn2.Close()
		if err != nil {
			log.Fatalf("Connection failed")
		}

		accept1 := getPeerConnection(conn, int64(s.timestamp))
		accept2 := getPeerConnection(conn2, int64(s.timestamp))
		if accept1 && accept2 {
			if in.Bid > int64(currentHighestBid) {
				currentHighestBid = int(in.Bid)
				fmt.Println("I am node ", s.id, " Current new price: ", currentHighestBid, " timestamp: ", s.timestamp)
			}
			fmt.Println(s.id, "finished asking for access")
			s.wantAccess = false
			return &proto.AcceptClientRequest{
				AuctionBid: int64(currentHighestBid),
				Giveacces:  true,
			}, nil

		} else {
			fmt.Println("I am ", s.id, " I got no access granted ", s.timestamp)
			time.Sleep(time.Millisecond * 100)
			s.timestamp += 5
		}
	}
}

func main() {
	server1 := &CommuncationServer{id: 1, otherServer1: "localhost:5052", otherServer2: "localhost:5053", timestamp: 1}
	server2 := &CommuncationServer{id: 2, otherServer1: "localhost:5051", otherServer2: "localhost:5053", timestamp: 2}
	server3 := &CommuncationServer{id: 3, otherServer1: "localhost:5051", otherServer2: "localhost:5052", timestamp: 3}

	Node1Add := ":5051"
	Node2Add := ":5052"
	Node3Add := ":5053"

	go server1.start_server(Node1Add)
	go server2.start_server(Node2Add)
	go server3.start_server(Node3Add)

	go server1.auction()
	go server2.auction()
	go server3.auction()

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
	s.status = "up"

	if err != nil {
		log.Fatalf("Did not work")
		s.status = "down"
	}

	err = grpcServer.Serve(listener)
}

func (s *CommuncationServer) auction() {
	for {
		time.Sleep(time.Millisecond * 100)

		if timeLeftOfAuction < 0 {
			fmt.Println("Auction should be over")
			break
		}
	}
}

func getPeerConnection(conn *grpc.ClientConn, timestamp int64) bool {
	client := proto.NewCommuncationClient(conn)

	accept, err := client.Request(context.Background(), &proto.RequestAccess{
		Time: timestamp,
	})

	if err != nil {
		log.Fatalf("Request failed")
	}

	if accept.Status == "up" {
		return accept.Giveacces
	} else {
		fmt.Println("why would server be down??")
		return accept.Giveacces
	}

}
