package raindropSearch

import (
	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

func RaindropSearch(keywords []string) ([]sqlite3.Raindrop, error) {
	raindrops, err := sqlite3.Search(keywords)
	if err != nil {
		return nil, err
	}

	return raindrops, nil
}
