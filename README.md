# hacker-news
Searches the top stories on [Hacker News](https://news.ycombinator.com) and outputs the title and link that match a regular expression.

### Installation

``` go get github.com/coyle/hacker-news```

### Dependencies

``` go get github.com/fatih/color```

### Run

```go run main.go -regexp="regexp to match on stories"```

Ex. ```go run main.go -regexp=[Nn]ode(.js)?```

*On a Mac you can CMD+click to open a URL from the terminal*
