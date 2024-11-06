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
	CriticalDataNumber = 0
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

	General1Add := ":5051"
	General2Add := ":5052"
	General3Add := ":5053"

	General1FullAdd := "localhost:5051"
	General2FullAdd := "localhost:5052"
	General3FullAdd := "localhost:5053"

	go server1.start_server(General1Add)
	go server2.start_server(General2Add)
	go server3.start_server(General3Add)

	go server1.connect(General2FullAdd, General3FullAdd)
	go server2.connect(General1FullAdd, General3FullAdd)
	go server3.connect(General1FullAdd, General2FullAdd)

	// Keep the main function alive to prevent exit
	select {}

}

func (s *CommuncationServer) start_server(GeneralAddress string) {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", GeneralAddress)
	if err != nil {
		log.Fatalf("Did not work")
	}

	proto.RegisterCommuncationServer(grpcServer, s)

	fmt.Println("opened server")

	if err != nil {
		log.Fatalf("Did not work")
	}

	err = grpcServer.Serve(listener)
}

func (s *CommuncationServer) connect(peer1 string, peer2 string) {
	s.timestamp = s.id

	for {

		wantAccessNumber := rand.Intn(3)

		if wantAccessNumber != 1 {
			for {
				s.wantAccess = true
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
					s.mu.Lock()
					CriticalDataNumber++
					s.mu.Unlock()
					fmt.Println("I am ", s.id, " Current number of Critical data: ", CriticalDataNumber, " timestamp: ", s.timestamp)
					time.Sleep(time.Millisecond * 100)
					break
				} else {
					fmt.Println("I am ", s.id, " I got no access granted ", s.timestamp)
					time.Sleep(time.Millisecond * 100)
					s.timestamp += 5
				}
			}
		} else {
			s.wantAccess = false
			time.Sleep(time.Millisecond * 100)
		}

	}
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
