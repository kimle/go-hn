package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Story struct
type Story struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Score       int    `json:"score"`
}

var flagvar int
var wg sync.WaitGroup

func init() {
	flag.IntVar(&flagvar, "stories", 0, "n stories to fetch and display")
	flag.Parse()
	if flagvar == 0 {
		color.Red("You must provide at least one story!")
		os.Exit(0)
	}
}

func printStories(stories []Story) {
	for i, story := range stories {
		c := color.New(color.FgCyan).Add(color.Underline)
		color.Cyan("%d.", i+1)
		c.Printf("%s [%d points]", story.Title, story.Score)
		//fmt.Printf("%d%s%d%s\n", i+1, ". "+story.Title+" [", story.Score, " points]")
		fmt.Printf("\t%d%s\t%s\n", story.Descendants, " comments", story.URL)
	}
}

func getIDsNew(n int, ch chan int) []Story {
	wg.Add(n)
	allIds := make([]int, 500)
	stories := make([]Story, n)
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader([]byte(string(body)))
	err = json.NewDecoder(r).Decode(&allIds)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < n; i++ {
		ch <- allIds[i]
		go getTopStoriesNew(i, ch, stories)
	}
	wg.Wait()
	close(ch)
	return stories
}

func getTopStoriesNew(idx int, ch <-chan int, stories []Story) {
	//now := time.Now()
	var story Story
	id := <-ch
	defer wg.Done()
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(id) + ".json")
	if err != nil {
		fmt.Println("Could not fetch story... Continuing...")
	}
	err = json.NewDecoder(resp.Body).Decode(&story)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Could not decode data... Continuing...")
	}
	fmt.Printf("%#v\n", story)
	stories[idx] = story
	//end := time.Now()
	//fmt.Printf("Loop %d took %f s\n", idx, end.Sub(now).Seconds())
}

func main() {
	start := time.Now()
	ch := make(chan int, flagvar)
	getIDsNew(flagvar, ch)
	t := time.Now()
	fmt.Printf("\n%s%f%s\n", "Time elapsed: ", t.Sub(start).Seconds(), "s")
}
