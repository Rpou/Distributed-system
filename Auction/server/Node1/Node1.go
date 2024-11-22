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

type CommuncationServer struct {
	proto.UnimplementedCommuncationServer
	timestamp         int
	id                int
	otherServer1      string
	otherServer2      string
	mu                sync.Mutex
	wantAccess        bool
	status            string
	mutex             sync.Mutex
	currentHighestBid int64
	timeLeftOfAuction int64
}

func (s *CommuncationServer) AuctionStatus(ctx context.Context, in *proto.Empty) (*proto.AuctionProgress, error) {
	return &proto.AuctionProgress{InProgress: s.timeLeftOfAuction > 0, HighestBid: int64(s.currentHighestBid)}, nil
}

func (s *CommuncationServer) Request(ctx context.Context, in *proto.RequestAccess) (*proto.AcceptNodeRequest, error) {

	// Accept if the other timestamp is smaller than this clients timestamp, or if this client does not want access

	return &proto.AcceptNodeRequest{Giveacces: in.Time > int64(s.timestamp) || !s.wantAccess, MyBid: int64(s.currentHighestBid), TimeLeftOfAuction: int64(s.timeLeftOfAuction)}, nil
}

func (s *CommuncationServer) ClientRequest(ctx context.Context, in *proto.ClientToNodeBid) (*proto.AcceptClientRequest, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	fmt.Println(s.id, "started asking for access")

	for {

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

		if accept1.Giveacces && accept2.Giveacces {
			auctionTime := min(accept1.TimeLeftOfAuction, accept2.TimeLeftOfAuction, s.timeLeftOfAuction)
			s.timeLeftOfAuction = auctionTime

			highestBid := max(s.currentHighestBid, accept1.MyBid, accept2.MyBid)
			s.currentHighestBid = highestBid

			if in.Bid > s.currentHighestBid {
				s.currentHighestBid = in.Bid
				fmt.Println("I am node ", s.id, " Current new price: ", s.currentHighestBid, " timestamp: ", s.timestamp)
				s.wantAccess = false
				return &proto.AcceptClientRequest{
					AuctionBid: s.currentHighestBid,
					Giveacces:  true,
				}, nil
			}
			fmt.Println(s.id, "finished asking for access")
			s.wantAccess = false
			return &proto.AcceptClientRequest{
				AuctionBid: s.currentHighestBid,
				Giveacces:  false,
			}, nil

		} else {
			fmt.Println("I am ", s.id, " I got no access granted ", s.timestamp)
			time.Sleep(time.Millisecond * 500)
			s.timestamp += 5
		}
	}
}

func main() {
	server1 := &CommuncationServer{id: 1, otherServer1: "localhost:5052", otherServer2: "localhost:5053", timestamp: 1, timeLeftOfAuction: 400}

	Node1Add := ":5051"

	go server1.start_server(Node1Add)

	go server1.auction()

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
		s.timeLeftOfAuction--
		if s.timeLeftOfAuction < 0 {
			fmt.Println("Auction should be over")

			conn, err := grpc.NewClient(s.otherServer1, grpc.WithTransportCredentials(insecure.NewCredentials()))
			conn2, err := grpc.NewClient(s.otherServer2, grpc.WithTransportCredentials(insecure.NewCredentials()))
			defer conn.Close()
			defer conn2.Close()
			accept1 := getPeerConnection(conn, int64(s.timestamp))
			accept2 := getPeerConnection(conn2, int64(s.timestamp))

			if err != nil {
				log.Fatalf("Connection failed")
			}

			highestBid := max(s.currentHighestBid, accept1.MyBid, accept2.MyBid)
			s.currentHighestBid = highestBid

			fmt.Println("Highest bid was", s.currentHighestBid)
			break
		}
	}
}

func getPeerConnection(conn *grpc.ClientConn, timestamp int64) *proto.AcceptNodeRequest {
	client := proto.NewCommuncationClient(conn)

	accept, err := client.Request(context.Background(), &proto.RequestAccess{
		Time: timestamp,
	})

	if err != nil {
		fmt.Println("Request failed")
		return &proto.AcceptNodeRequest{
			Giveacces:         true,
			MyBid:             0,
			TimeLeftOfAuction: 1000,
		}
	}

	return accept

}
