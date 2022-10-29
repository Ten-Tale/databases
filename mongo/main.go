package main

import (
	"context"
	"databases/mongo/databaseworker"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client := databaseworker.ConnectToDatabase()
	defer databaseworker.CloseDbConnection(client)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := client.Ping(ctx, readpref.Primary())

	checkError(err)

	collection := client.Database("group").Collection("students")

	fmt.Println("here")

	// INSERT
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{
		{Key: "id", Value: "19B0544"},
		{Key: "firstName", Value: "Andrey"},
		{Key: "lastName", Value: "Shkunov"},
	})

	checkError(err)

	fmt.Println(res.InsertedID)

	// GET
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	foundValue := collection.FindOne(ctx, bson.D{{Key: "id", Value: "19B0544"}})

	fmt.Println(foundValue.DecodeBytes())

	// UPDATE
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection.UpdateOne(ctx,
		bson.D{{Key: "id", Value: "19B0544"}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "age", Value: 21},
		},
		}},
	)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	foundValue = collection.FindOne(ctx, bson.D{{Key: "id", Value: "19B0544"}})

	fmt.Println(foundValue.DecodeBytes())

	// DELETE
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"id": "19B0544"})

	checkError(err)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	foundValue = collection.FindOne(ctx, bson.D{{Key: "id", Value: "19B0544"}})

	fmt.Println(foundValue.DecodeBytes())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
