package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"URL_Shortener/db"

	"github.com/redis/go-redis/v9"
)

type URL struct {
	ID    int    `json:"id"`
	Full  string `json:"full"`
	Short string `json:"short"`
}

type RedisCli struct {
	Conn *redis.Client
}

func RDB(d *db.DBshort) (*RedisCli, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	rows, err := d.DB.Query("SELECT Id, FullURL, ShortURL FROM urls")
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u URL
		if err := rows.Scan(&u.ID, &u.Full, &u.Short); err != nil {
			log.Println("row scan error:", err)
			continue
		}

		data, err := json.Marshal(u)
		if err != nil {
			log.Println("json marshal error:", err)
			continue
		}

		key := fmt.Sprintf("url:%d", u.ID)
		if err := rdb.Set(ctx, key, data, 1*time.Hour).Err(); err != nil {
			log.Println("redis set error:", err)
			continue
		}
	}

	return &RedisCli{Conn: rdb}, nil
}

func (rdb *RedisCli) GetData() {
	ctx := context.Background()
	data, err := rdb.Conn.Get(ctx, "url:1").Result()
	if err != nil {
		log.Println("redis get error:", err)
		return
	}

	log.Println("Data from Redis:", data)
}
