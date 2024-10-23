package main

import (
	proto "ChittyChat/grpc"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	"google.golang.org/grpc"
)

type ChittychatDBServer struct {
	proto.UnimplementedChittychatDBServer
	posts             []string
	serverLamportTime int64
	mu                sync.Mutex
}

func (s *ChittychatDBServer) GetPosts(ctx context.Context, in *proto.ClientLT) (*proto.Posts, error) {
	s.serverLamportTime = max(s.serverLamportTime, in.LamportTime) + 1

	return &proto.Posts{Posts: s.posts, LamportTime: s.serverLamportTime}, nil
}

func (s *ChittychatDBServer) Connect(ctx context.Context, in *proto.ClientInfo) (*proto.Empty, error) {
	s.serverLamportTime = max(s.serverLamportTime, in.LamportTime) + 1
	messageWithLamportTime := " Lamport time: (" + strconv.FormatInt(s.serverLamportTime, 10) + ", " + strconv.FormatInt(in.Cn, 10) + ")"
	s.posts = append(s.posts, fmt.Sprintf("The following client connected: %d"+messageWithLamportTime, in.Cn))
	return &proto.Empty{}, nil
}

func (s *ChittychatDBServer) Disconnect(ctx context.Context, in *proto.ClientInfo) (*proto.Empty, error) {
	s.serverLamportTime = max(s.serverLamportTime, in.LamportTime) + 1
	messageWithLamportTime := " Lamport time: (" + strconv.FormatInt(s.serverLamportTime, 10) + ", " + strconv.FormatInt(in.Cn, 10) + ")"
	s.posts = append(s.posts, fmt.Sprintf("The following client disconnected: %d"+messageWithLamportTime, in.Cn))
	return &proto.Empty{}, nil
}

func (s *ChittychatDBServer) PublishPost(ctx context.Context, in *proto.Post) (*proto.Posted, error) {
	//log.Printf("Received post: %s with Lamport time: %d", in.Post, in.LamportTime)
	s.serverLamportTime = max(s.serverLamportTime, in.LamportTime) + 1

	if len(in.Post) <= 128 {
		postWithLamportTime := in.Post + ", Lamport time: (" + strconv.FormatInt(s.serverLamportTime, 10) + ", " + string(in.Post[len(in.Post)-1]) + ")"
		s.posts = append(s.posts, postWithLamportTime)
		return &proto.Posted{Posted: true, LamportTime: s.serverLamportTime}, nil
	}
	return &proto.Posted{Posted: false, LamportTime: s.serverLamportTime}, nil
}

func main() {
	server := &ChittychatDBServer{posts: []string{}, serverLamportTime: 1}
	server.posts = append(server.posts, "We have begun")

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
