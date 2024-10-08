package main

import (
	proto "ChittyChat/grpc"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChittychatDBServer struct {
	proto.UnimplementedChittychatDBServer
	posts []string
}

func (s *ChittychatDBServer) GetPosts(ctx context.Context, in *proto.Empty) (*proto.Posts, error) {
	return &proto.Posts{Posts: s.posts}, nil
}

func main() {
	server := &ChittychatDBServer{posts: []string{}}
	server.posts = append(server.posts, "ma dick big")

	server.start_server()
}

func (s *ChittychatDBServer) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatalf("Did not work")
	}

	proto.RegisterChittychatDBServer(grpcServer, s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("du en hoe")
	}

}
