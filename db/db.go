package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DBshort struct {
	DB *sql.DB
}

func Conn() (*DBshort, error) {
	connStr := "user=ilya password=123 dbname=shortener host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DBshort{DB: db}, nil
}

func (d *DBshort) Add(fullUrl, shortUrl string) error {
	query := "INSERT INTO urls (FullURL, ShortURL) VALUES ($1, $2)"
	_, err := d.DB.Exec(query, fullUrl, shortUrl)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBshort) Get(shortUrl string) (string, error) {
	var full string

	err := d.DB.QueryRow(
		"SELECT FullURL FROM urls WHERE ShortURL = $1",
		shortUrl,
	).Scan(&full)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
		log.Println(err)
		return "", err
	}

	return full, nil
}
