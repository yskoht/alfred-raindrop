package main

import (
	"os"

	raindropIncrement "github.com/yskoht/alfred-raindrop/internal/raindrop-increment"
	"github.com/yskoht/alfred-raindrop/pkg/file"
	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

func run() {
	if !file.Exists(sqlite3.DB_FILE) {
		return
	}

	id := os.Args[1]

	err := raindropIncrement.RaindropIncrementViewCount(id)
	if err != nil {
		return
	}
}

func main() {
	run()
}
