package raindropSync

import (
	"github.com/yskoht/alfred-raindrop/pkg/raindrop"
	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

func RaindropSync(token string) error {
	raindrops, err := raindrop.GetAllRaindrops(token)
	if err != nil {
		return err
	}

	err = sqlite3.CreateDB(raindrops)
	if err != nil {
		return err
	}

	return nil
}
