package main

import (
	"fmt"
	"os"

	raindropFind "github.com/yskoht/alfred-raindrop/internal/raindrop-find"
	"github.com/yskoht/alfred-raindrop/pkg/file"
	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

func run() {
	if !file.Exists(sqlite3.DB_FILE) {
		return
	}

	id := os.Args[1]

	raindrop, err := raindropFind.RaindropFind(id)
	if err != nil {
		return
	}

	fmt.Print(raindrop.Link)
}

func main() {
	run()
}
