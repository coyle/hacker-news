package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	"github.com/fatih/color"
)

var (
	base           = "https://hacker-news.firebaseio.com"
	item           = base + "/v0/item/"
	stories        = base + "/v0/topstories.json"
	storiesToMatch = " [Gg][Oo](lang)? | [Nn]ode(.js)?"
	wg             sync.WaitGroup
	link           = color.New(color.FgCyan).Add(color.Underline)
	title          = color.New(color.FgWhite)
	mutex          = &sync.Mutex{}
)

type story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	resp, err := http.Get(stories)
	if err != nil {
		fmt.Println(err)
		return
	}

	var top []int
	err = json.NewDecoder(resp.Body).Decode(&top)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range top {
		go getStory(val)
	}
	wg.Wait()
	fmt.Println()
}

func getStory(stry int) {
	wg.Add(1)
	defer wg.Done()
	resp, err := http.Get(item + strconv.Itoa(stry) + ".json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var story story
	err = json.NewDecoder(resp.Body).Decode(&story)
	if err != nil {

	}
	matchStory(story)
}

func matchStory(story story) {
	if matched, _ := regexp.MatchString(storiesToMatch, story.Title); matched {
		mutex.Lock()
		title.Printf("\n\t%s\n", story.Title)
		link.Printf("\t%s\n", story.URL)
		mutex.Unlock()
	}
}
