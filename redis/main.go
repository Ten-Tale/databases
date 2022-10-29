package main

import (
	"context"
	"databases/redis/databaseworker"
	"fmt"
)

func main() {
	ctx := context.Background()

	rdb := databaseworker.ConnectToDB()
	defer rdb.Close()

	rdb.HSet(ctx, "administrator",
		"id", "19B0544",
		"firstname", "Andrey",
		"lastname", "Shkunov",
	)

	result := rdb.HGet(ctx, "administrator", "id")

	fmt.Println(result.Result())

	rdb.Del(ctx, "administrator")

	resAll := rdb.HGetAll(ctx, "administrator")

	fmt.Println(resAll.Result())
}
