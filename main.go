// hacker-news searches the top stories on https://news.ycombinator.com and outputs the title and link that match a regex
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	base           = "https://hacker-news.firebaseio.com"
	item           = base + "/v0/item/"
	stories        = base + "/v0/topstories.json"
	wg             sync.WaitGroup
	link           = color.New(color.FgCyan).Add(color.Underline)
	title          = color.New(color.FgWhite)
	bg             = color.New(color.BgHiYellow)
	mutex          = &sync.Mutex{}
	defaultRegexp  = ".*"
	cnt            = 0
	storiesToMatch string
)

type story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	flag.StringVar(&storiesToMatch, "regexp", defaultRegexp, "The regular expression to match Hacker News stories on")
	flag.Parse()

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
	resp, err := http.Get(buildStoryURL(stry))
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
		if story.URL == "" {
			return
		}
		mutex.Lock()
		cnt++
		title.Printf("\n%d.\t%s\n", cnt, story.Title)
		link.Printf("\t%s\n", story.URL)
		title.Printf("\n%s\n\n", strings.Repeat("_", 100))
		mutex.Unlock()
	}
}

func buildStoryURL(stry int) string {
	return item + strconv.Itoa(stry) + ".json"
}
