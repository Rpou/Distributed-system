package main

import (
	proto "ITUServer/grpc"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CommuncationServer struct {
	proto.UnimplementedCommuncationServer
	timestamp int64
}

func (s *CommuncationServer) Request(ctx context.Context, in *proto.CriticalData) (*proto.Accept, error) {
	fmt.Println(in.CriticalData)
	return &proto.Accept{Giveacces: true}, nil
}

func main() {
	server := &CommuncationServer{}
	General1Add := ":5051"
	General2Add := ":5052"
	General3Add := ":5053"

	go server.start_server(General1Add)
	go server.start_server(General2Add)
	go server.start_server(General3Add)
	go server.connect(General1Add)
	go server.connect(General2Add)
	go server.connect(General3Add)

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

func (s *CommuncationServer) connect(GeneralAddress string) {
	for {
		time.Sleep(time.Second)
		wantAccess := rand.Intn(3)
		timestamp := rand.Intn(20)

		if wantAccess == 0 && GeneralAddress == ":5053" {
			conn, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			conn2, err := grpc.NewClient("localhost:5052", grpc.WithTransportCredentials(insecure.NewCredentials()))
			defer conn.Close()
			if err != nil {
				log.Fatalf("Did not work")
			}

			client := proto.NewCommuncationClient(conn)

			fmt.Println("connected first")

			accept, err := client.Request(context.Background(), &proto.CriticalData{
				CriticalData: 2,
				Time:         int64(timestamp),
			})

			if accept.Giveacces {
				fmt.Println("Access granted from client 1")
			} else {
				fmt.Println("No acces granted from client 1")
			}

			client = proto.NewCommuncationClient(conn2)

			accept, err = client.Request(context.Background(), &proto.CriticalData{
				CriticalData: 2,
				Time:         int64(timestamp),
			})

			if accept.Giveacces {
				fmt.Println("Access granted from client 2")
			} else {
				fmt.Println("No acces granted from client 2")
			}

		}

	}
}
