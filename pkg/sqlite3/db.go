package sqlite3

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yskoht/alfred-raindrop/pkg/file"
	"github.com/yskoht/alfred-raindrop/pkg/raindrop"
)

const (
	DB_FILE = "./db/db.sqlite3"
)

func RemoveDB() error {
	if !file.Exists(DB_FILE) {
		return nil
	}

	return os.Remove(DB_FILE)
}

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
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return err
	}
	defer db.Close()

	dropTable := "drop table if exists raindrops"
	_, err = db.Exec(dropTable)
	if err != nil {
		return err
	}

	createRaindropsTable :=
		"create table raindrops (" +
			"id integer not null primary key," +
			"title text not null," +
			"link text not null" +
			")"
	_, err = db.Exec(createRaindropsTable)
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

	createViewCountsTable :=
		"create table" +
			" if not exists view_counts (" +
			" id integer not null primary key autoincrement," +
			" raindrop_id integer not null unique," +
			" count integer not null" +
			")"
	_, err = db.Exec(createViewCountsTable)
	if err != nil {
		return err
	}

	return nil
}

type Raindrop struct {
	ID    int
	Title string
	Link  string
}

func convertRowToRaindrop(row *sql.Row) (*Raindrop, error) {
	var id int
	var title string
	var link string

	err := row.Scan(&id, &title, &link)
	if err != nil {
		return nil, err
	}

	return &Raindrop{ID: id, Title: title, Link: link}, nil
}

func convertRowsToRaindrop(rows *sql.Rows) (*Raindrop, error) {
	var id int
	var title string
	var link string

	err := rows.Scan(&id, &title, &link)
	if err != nil {
		return nil, err
	}

	return &Raindrop{ID: id, Title: title, Link: link}, nil
}

func Search(keywords []string) ([]Raindrop, error) {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	search, err := db.Prepare(
		"select raindrops.id, raindrops.title, raindrops.link" +
			" from raindrops" +
			" left outer join view_counts on raindrops.id = view_counts.raindrop_id" +
			" where title like ? or link like ?" +
			" order by view_counts.count desc",
	)

	if err != nil {
		return nil, err
	}

	pattern := "%" + strings.Join(keywords, "%") + "%"
	rows, err := search.Query(pattern, pattern)
	if err != nil {
		return nil, err
	}

	raindrops := make([]Raindrop, 0)
	for rows.Next() {
		raindrop, err := convertRowsToRaindrop(rows)
		if err != nil {
			return nil, err
		}
		raindrops = append(raindrops, *raindrop)
	}

	return raindrops, nil
}

func Find(id int) (*Raindrop, error) {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(
		"select id, title, link from raindrops where id = ?",
	)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)
	raindrop, err := convertRowToRaindrop(row)
	if err != nil {
		return nil, err
	}

	return raindrop, nil
}

func IncrementViewCount(id int) error {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(
		"insert into view_counts(raindrop_id, count) values (?, 1)" +
			" on conflict(raindrop_id)" +
			" do update set count = count + 1",
	)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
