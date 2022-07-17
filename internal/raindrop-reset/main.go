package raindropReset

import (
	"github.com/yskoht/alfred-raindrop/pkg/sqlite3"
)

func RaindropReset() error {
	return sqlite3.RemoveDB()
}
