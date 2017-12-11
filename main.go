package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

func init() {
	flag.IntVar(&flagvar, "stories", 0, "n stories to fetch and display")
	flag.Parse()
	if flagvar == 0 {
		color.Red("You must provide at least one story!")
		os.Exit(0)
	}
}

func getIDs(n int) []int {
	red := color.New(color.FgRed).Add(color.Bold)
	allIds := make([]int, 500)
	idList := make([]int, n)
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		red.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		red.Println(err)
		os.Exit(1)
	}

	r := bytes.NewReader([]byte(string(body)))
	err = json.NewDecoder(r).Decode(&allIds)
	if err != nil {
		red.Println(err)
		os.Exit(1)
	}

	for i := 0; i < n; i++ {
		idList[i] = allIds[i]
	}
	return idList
}

func getTopStories(IDs []int) []Story {
	var story Story
	topStories := make([]Story, len(IDs))
	red := color.New(color.FgRed).PrintfFunc()

	for i := 0; i < len(IDs); i++ {
		resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(IDs[i]) + ".json")
		if err != nil {
			red("Could not fetch story... Continuing...")
			continue
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&story)
		if err != nil {
			red("Could not decode data... Continuing...")
			continue
		}
		topStories[i] = story
	}
	return topStories
}

func printStories(stories []Story) {
	cyan := color.New(color.FgHiCyan).Add(color.Bold)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen).Add(color.Underline)
	for i, story := range stories {
		cyan.Printf("%d%s", i+1, ". ")
		color.Magenta("%s [%d %s]", story.Title, story.Score, "points")
		yellow.Printf("\t%d %s", story.Descendants, "comments")
		green.Printf("\t%s\n", story.URL)
	}
}

func main() {
	fmt.Print("Starting timer...\n\n")
	start := time.Now()

	IDs := getIDs(flagvar)
	stories := getTopStories(IDs)
	printStories(stories)

	t := time.Now().Sub(start).Round(time.Millisecond).Truncate(time.Millisecond).String()
	fmt.Printf("\n%s%s\n", "Time elapsed (in seconds): ", t)
}
