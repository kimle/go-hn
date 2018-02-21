package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/kimle/go-hackernews/service"
	"google.golang.org/grpc"
)

const (
	// address = "localhost:50051"
	address = "35.229.113.24:50051"
)

func main() {
	allIds := make([]int32, 500)
	stories := make([]*pb.Story, 10)
	resp, _ := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response := bytes.NewReader([]byte(string(body)))
	json.NewDecoder(response).Decode(&allIds)
	for i := 0; i < 10; i++ {
		stories[i] = &pb.Story{Id: allIds[i]}
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTestClient(conn)
	log.Printf("argument: %v", &pb.TopStories{TopStories: stories})
	r, err := c.GetStories(context.Background(), &pb.TopStories{TopStories: stories})
	log.Printf("%v", r)
}
