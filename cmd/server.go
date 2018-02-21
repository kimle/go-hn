package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"

	pb "github.com/kimle/go-hackernews/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetStory(ctx context.Context, in *pb.TopStories) (*pb.Story, error) {
	log.Printf("TopStories: %v\n", in)
	return &pb.Story{Id: in.TopStories[0].GetId()}, nil
}

func (s *server) GetStories(ctx context.Context, in *pb.TopStories) (*pb.Stories, error) {
	stories := make([]*pb.Story, 10)
	for i := 0; i < 10; i++ {
		resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" +
			strconv.FormatInt(int64(in.TopStories[i].GetId()), 10) + ".json")
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		err = json.NewDecoder(resp.Body).Decode(&stories[i])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("story %d: %v", i, stories[i])
		defer resp.Body.Close()
	}
	return &pb.Stories{stories}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on port: %v", port)
	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
