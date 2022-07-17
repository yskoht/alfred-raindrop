package main

import (
	aw "github.com/deanishe/awgo"
	raindropSearch "github.com/yskoht/alfred-raindrop/internal/raindrop-search"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	keywords := wf.Args()

	raindrops, err := raindropSearch.RaindropSearch(keywords)
	if err != nil {
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
