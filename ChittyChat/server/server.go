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

/*
	func (s *ChittychatDBServer) Connect(ctx context.Context, in *proto.ClientNumber) (*proto.Connected, error) {
		return &proto.Connected{Con: true}, nil
	}

	func (s *ChittychatDBServer) Disconnect(ctx context.Context, in *proto.ClientNumber) (*proto.Connected, error) {
		return &proto.Connected{Con: true}, nil
	}
*/

func (s *ChittychatDBServer) PublishPost(ctx context.Context, in *proto.Post) (*proto.Posted, error) {
	//log.Printf("Received post: %s with Lamport time: %d", in.Post, in.LamportTime)
	if len(in.Post) <= 128 {
		s.posts = append(s.posts, in.Post)
		return &proto.Posted{Posted: true}, nil
	}
	return &proto.Posted{Posted: false}, nil
}

func main() {
	server := &ChittychatDBServer{posts: []string{}}
	server.posts = append(server.posts, "Cause i got first like")

	server.start_server()
}

func (s *ChittychatDBServer) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatalf("Server could not start1")
	}

	proto.RegisterChittychatDBServer(grpcServer, s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Server could not start2")
	}

}
