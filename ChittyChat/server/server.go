package main

import (
	proto "ChittyChat/grpc"
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type ChittychatDBServer struct {
	proto.UnimplementedChittychatDBServer
	posts             []string
	serverLamportTime int64
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
	if in.LamportTime > int64(s.serverLamportTime) {
		s.serverLamportTime = int64(in.LamportTime) + 1
	} else {
		s.serverLamportTime += 1
	}

	if len(in.Post) <= 128 {
		postWithLamportTime := in.Post + " " + strconv.FormatInt(in.LamportTime, 10)
		s.posts = append(s.posts, postWithLamportTime)
		return &proto.Posted{Posted: true}, nil
	}
	return &proto.Posted{Posted: false}, nil
}

func main() {
	server := &ChittychatDBServer{posts: []string{}, serverLamportTime: 1}
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
