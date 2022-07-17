package raindropFind

import (
	"strconv"

	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

func RaindropFind(id string) (*sqlite3.Raindrop, error) {
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	raindrop, err := sqlite3.Find(_id)
	if err != nil {
		return nil, err
	}

	return raindrop, nil
}
