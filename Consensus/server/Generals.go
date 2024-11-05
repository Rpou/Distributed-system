package main

import (
	proto "ITUServer/grpc"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CommuncationServer struct {
	proto.UnimplementedCommuncationServer
}

func (s *CommuncationServer) Request(ctx context.Context, in *proto.CriticalData) (*proto.Accept, error) {
	fmt.Print(in.CriticalData)
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
	if GeneralAddress == "5053" {
		conn, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		defer conn.Close()
		if err != nil {
			log.Fatalf("Did not work")
		}
		client := proto.NewCommuncationClient(conn)

		accept, err := client.Request(context.Background(), &proto.CriticalData{
			CriticalData: 2,
			Time:         90,
		})
		if accept.Giveacces {
			fmt.Println(".fasd")
		}
	}
	err = grpcServer.Serve(listener)

}
