package raindropIncrement

import (
	"strconv"

	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

func RaindropIncrementViewCount(id string) error {
	_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = sqlite3.IncrementViewCount(_id)
	if err != nil {
		return err
	}

	return nil
}
