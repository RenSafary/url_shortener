package redis

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"URL_Shortener/db"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type url struct {
	ID    int    `json:"id"`
	Full  string `json:"full"`
	Short string `json:"short"`
}

func RDB(d *db.DBshort) error {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	rows, err := d.DB.Query("SELECT * FROM urls")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
			return nil
		} else {
			log.Fatal(err)
			return nil
		}
	}

	for rows.Next() {
		var id int
		var fullUrl, shortUrl string

		if err := rows.Scan(&id, &fullUrl, &shortUrl); err != nil {
			log.Fatal(err)
			return err
		}

		data, err := json_marshal(id, fullUrl, shortUrl)
		if err != nil {
			log.Fatal(err)
			return err
		}

		key := fmt.Sprintf("url:%d", id)
		err = rdb.Set(ctx, key, data, 10*time.Second).Err()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func json_marshal(id int, full, short string) ([]byte, error) {
	urlData := map[string]interface{}{
		"id":    id,
		"full":  full,
		"short": short,
	}

	return json.Marshal(urlData)
}
