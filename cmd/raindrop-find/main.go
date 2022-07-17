package main

import (
	raindropFind "github.com/yskoht/alfred-raindrop/internal/raindrop-find"
	"github.com/yskoht/alfred-raindrop/pkg/file"
	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	if !file.Exists(sqlite3.DB_FILE) {
		wf.Warn("Database not found", "Please run `raindrop-sync`")
		return
	}

	id := wf.Args()[0]

	raindrop, err := raindropFind.RaindropFind(id)
	if err != nil {
		wf.FatalError(err)
		return
	}

	wf.
		NewItem(raindrop.Title).
		Subtitle(raindrop.Link).
		Arg(raindrop.Link).
		Valid(true)

	wf.WarnEmpty("Bookmark not found", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
