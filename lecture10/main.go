package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var redisClient *redis.Client
var db *sql.DB

func main() {
	key := "hash"
	ctx := context.Background()
	var dataFromPostgre string
	cachedData, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			err := db.QueryRowContext(ctx, "SELECT data FROM lecture10 WHERE key = $1", key).Scan(&dataFromPostgre)
			if err != nil {
				log.Fatalf("error getting data from postgre %v", err)
			}
			err = redisClient.Set(ctx, key, dataFromPostgre, 10*time.Minute).Err()
			if err != nil {
				log.Fatalf("error setting data to redis %v", err)
			}
		} else {
			log.Fatalf("error getting data from redis %v", err)
		}
		fmt.Printf("From Postgre: %s", dataFromPostgre)

	} else {
		fmt.Printf("From Redis: %s", cachedData)
	}

}

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=qwerty dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
