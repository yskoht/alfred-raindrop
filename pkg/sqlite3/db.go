package sqlite3

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yskoht/alfred-raindrop/pkg/raindrop"
	"os"
	"strings"
)

const (
	DB_FILE    = "./db/db.sqlite3"
	TABLE_NAME = "raindrops"
	SCHEMA     = "id integer not null primary key, title text, link text"
	COLUMNS    = "id, title, link"
)

func insertValues(raindrops []raindrop.Raindrop) string {
	values := make([]string, 0)
	for _, raindrop := range raindrops {
		title := strings.Replace(raindrop.Title, "'", "", -1)
		link := strings.Replace(raindrop.Link, "'", "", -1)
		value := fmt.Sprintf("(%d, '%s', '%s')", raindrop.ID, title, link)
		values = append(values, value)
	}
	return strings.Join(values, ",")
}

func CreateDB(raindrops []raindrop.Raindrop) error {
	os.Remove(DB_FILE)

	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return err
	}
	defer db.Close()

	createTable := fmt.Sprintf("%s %s (%s);", "create table", TABLE_NAME, SCHEMA)
	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	if len(raindrops) > 0 {
		insertValues := insertValues(raindrops)
		insert := fmt.Sprintf("%s %s (%s) values %s", "insert into", TABLE_NAME, COLUMNS, insertValues)
		_, err = db.Exec(insert)
		if err != nil {
			return err
		}
	}

	return nil
}
