package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/kimle/go-hackernews/service"
	"google.golang.org/grpc"
)

var flagvar int
var env string
var address string

func init() {
	flag.IntVar(&flagvar, "stories", 0, "n stories to fetch and display")
	flag.StringVar(&env, "env", "dev", "enviroment to use")
	flag.Parse()
	if flagvar == 0 {
		fmt.Println("you must provide at least one story!")
		os.Exit(0)
	}

	if env == "dev" {
		address = "localhost:50051"
	} else if env == "prod" {
		address = "35.229.113.24:50051"
	}
}

func main() {
	stories := make([]*pb.Story, flagvar)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	c := pb.NewTestClient(conn)
	r, err := c.GetIds(ctx, &pb.Amount{Amount: int32(flagvar)})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < flagvar; i++ {
		stories[i] = &pb.Story{Id: r.Ids[i]}
	}
	rStories, err := c.GetStories(ctx, &pb.TopStories{TopStories: stories})
	for i, story := range rStories.Stories {
		fmt.Printf("\n%d. %v\n url: %v\n", i+1, story.Title, story.Url)
	}
}
