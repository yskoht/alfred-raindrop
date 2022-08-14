package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	aw "github.com/deanishe/awgo"

	raindropFind "github.com/yskoht/alfred-raindrop/internal/raindrop-find"
	raindropIncrement "github.com/yskoht/alfred-raindrop/internal/raindrop-increment"
	raindropReset "github.com/yskoht/alfred-raindrop/internal/raindrop-reset"
	raindropSearch "github.com/yskoht/alfred-raindrop/internal/raindrop-search"
	raindropSync "github.com/yskoht/alfred-raindrop/internal/raindrop-sync"
	"github.com/yskoht/alfred-raindrop/pkg/file"
	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

var wf *aw.Workflow
var args []string

func init() {
	wf = aw.New()
}

func getToken() (string, error) {
	testToken := os.Getenv("testToken")
	if testToken == "" {
		return "", errors.New("testToken not found")
	}

	token := fmt.Sprintf("Bearer %s", testToken)
	return token, nil
}

func sync() {
	token, err := getToken()
	if err != nil {
		fmt.Print(err)
		return
	}

	err = raindropSync.RaindropSync(token)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func search() {
	if !file.Exists(sqlite3.DB_FILE) {
		wf.Warn("Database not found", "Please run `raindrop-sync`")
		return
	}

	keywords := args

	raindrops, err := raindropSearch.RaindropSearch(keywords)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, raindrop := range raindrops {
		wf.
			NewItem(raindrop.Title).
			Subtitle(raindrop.Link).
			Arg(strconv.Itoa(raindrop.ID)).
			Valid(true)
	}

	wf.WarnEmpty("Bookmark not found", "Try a different query?")
	wf.SendFeedback()
}

func reset() {
	err := raindropReset.RaindropReset()
	if err != nil {
		fmt.Print(err)
		return
	}
}

func increment() {
	if !file.Exists(sqlite3.DB_FILE) {
		return
	}

	id := args[0]

	err := raindropIncrement.RaindropIncrementViewCount(id)
	if err != nil {
		return
	}
}

func find() {
	if !file.Exists(sqlite3.DB_FILE) {
		return
	}

	id := args[0]

	raindrop, err := raindropFind.RaindropFind(id)
	if err != nil {
		return
	}

	fmt.Print(raindrop.Link)
}

func run() {
	if len(os.Args) < 2 {
		fmt.Println("hoge")
		os.Exit(1)
	}

	subCommand := os.Args[1]
	args = os.Args[2:]

	switch subCommand {
	case "sync":
		sync()
	case "search":
		wf.Run(search)
	case "reset":
		reset()
	case "increment":
		increment()
	case "find":
		find()
	}
}

func main() {
	run()
}
