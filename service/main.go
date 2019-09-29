package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wandi34/decision-helper/service/decision_helper"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var topics = []string{"Tribe Name", "App Name"}
var decisions1 = []string{"TribeClub", "Space Taxi"}
var decisions2 = []string{"Tinder", "DecisionFinder"}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetTopics(ctx context.Context, in *pb.TopicsRequest) (*pb.TopicsReply, error) {
	log.Printf("Received GetTopicsRequest:")
	return &pb.TopicsReply{Topics: topics}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetDecisionsForTopic(ctx context.Context, in *pb.DecisionsRequest) (*pb.DecisionsReply, error) {
	log.Printf("Received GetDecisionsRequest:")
	if in.GetTopic() == "TribeName" {
		return &pb.DecisionsReply{Decisions: decisions1}, nil
	}
	return &pb.DecisionsReply{Decisions: decisions2}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDecisionHelperServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
