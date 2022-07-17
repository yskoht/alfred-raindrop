package main

import (
	"os"

	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"

	aw "github.com/deanishe/awgo"
	raindropSearch "github.com/yskoht/alfred-raindrop/internal/raindrop-search"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func run() {
	if !exists(sqlite3.DB_FILE) {
		wf.Warn("Database not found", "Please type `raindrop-sync`")
		return
	}

	keywords := wf.Args()

	raindrops, err := raindropSearch.RaindropSearch(keywords)
	if err != nil {
		wf.FatalError(err)
		return
	}

	for _, raindrop := range raindrops {
		wf.
			NewItem(raindrop.Title).
			Subtitle(raindrop.Link).
			Arg(raindrop.Link).
			Valid(true)
	}

	wf.WarnEmpty("Bookmark not found", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
