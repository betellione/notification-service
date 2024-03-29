package main

import (
	"context"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "notification-service/proto"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *server) SendErrorNotification(ctx context.Context, in *pb.ErrorNotification) (*pb.NotificationResponse, error) {
	err := sendEmail(in.ErrorMessage)
	if err != nil {
		return &pb.NotificationResponse{Success: false}, err
	}
	return &pb.NotificationResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
