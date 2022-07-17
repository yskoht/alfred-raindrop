package sqlite3

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yskoht/alfred-raindrop/pkg/raindrop"
)

const (
	DB_FILE = "./db/db.sqlite3"
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

	createTable := "create table raindrops (id integer not null primary key, title text, link text)"
	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	if len(raindrops) > 0 {
		insertValues := insertValues(raindrops)
		insert := fmt.Sprintf(
			"insert into raindrops (id, title, link) values %s", insertValues,
		)
		_, err = db.Exec(insert)
		if err != nil {
			return err
		}
	}

	return nil
}

type Raindrop struct {
	ID    int
	Title string
	Link  string
}

func Search(keywords []string) ([]Raindrop, error) {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	pattern := "'%" + strings.Join(keywords, "%") + "%'"
	search := fmt.Sprintf(
		"select id, title, link from raindrops where title like %s or link like %s", pattern, pattern,
	)
	rows, err := db.Query(search)
	if err != nil {
		return nil, err
	}

	raindrops := make([]Raindrop, 0)
	for rows.Next() {
		var id int
		var title string
		var link string
		err = rows.Scan(&id, &title, &link)
		if err != nil {
			return nil, err
		}
		raindrops = append(raindrops, Raindrop{ID: id, Title: title, Link: link})
	}

	return raindrops, nil
}
