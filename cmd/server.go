package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (s *server) GetIds(ctx context.Context, in *pb.Amount) (*pb.Ids, error) {
	if in.Amount > 500 {
		log.Fatalf("cannot get more than 500 stories")
	}
	allIds := make([]int32, 500)
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatalf("could not get ids: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := bytes.NewReader([]byte(string(body)))
	err = json.NewDecoder(response).Decode(&allIds)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.Ids{Ids: allIds[:in.Amount]}, nil
}

func (s *server) GetStory(ctx context.Context, in *pb.TopStories) (*pb.Story, error) {
	log.Printf("TopStories: %v\n", in)
	return &pb.Story{Id: in.TopStories[0].GetId()}, nil
}

func (s *server) GetStories(ctx context.Context, in *pb.TopStories) (*pb.Stories, error) {
	amount := len(in.TopStories)
	fmt.Printf("amount: %v")
	stories := make([]*pb.Story, amount)
	for i := 0; i < amount; i++ {
		resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" +
			strconv.FormatInt(int64(in.TopStories[i].GetId()), 10) + ".json")
		if err != nil {
			log.Fatal(err)
		}
		err = json.NewDecoder(resp.Body).Decode(&stories[i])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("story %d: %v", i, stories[i])
		defer resp.Body.Close()
	}
	fmt.Printf("stories: %v", stories)
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
