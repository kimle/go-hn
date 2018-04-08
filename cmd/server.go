package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	pb "github.com/kimle/go-hackernews/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

var wg sync.WaitGroup

type server struct{}

func (s *server) GetIds(ctx context.Context, in *pb.Amount) (*pb.Ids, error) {
	if in.Amount > 500 {
		log.Fatalf("cannot get more than 500 stories")
	}
	allIds := make([]int32, in.Amount)
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatalf("could not get ids: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &allIds)
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
	stories := make([]*pb.Story, amount)
	for i := 0; i < amount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
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
		}(i)
	}
	wg.Wait()
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
