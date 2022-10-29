package databaseworker

import (
	"github.com/go-redis/redis/v9"
)

func ConnectToDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB

	})

}
