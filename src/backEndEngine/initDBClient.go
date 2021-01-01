package main

import (
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
)

// function to open up a mongo conection
func initDB() (*mongo.Client, *context.Context, error) {

	// mongodb://myDBReader:D1fficultP%40ssw0rd@mongodb0.example.com:27017/?authSource=admin
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_DB"))

	printInfo(fmt.Sprintf("Using Connection String [%s]", connectionString))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	// since Connect() is non-blocking at this stage, lets test to see if we actually got that last bit right
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		printError(fmt.Sprintf("Error Initialising the DB connection - %s", err))
		return nil, nil, err
	}

	printSuccess("Database found and connection established")

	return client, &ctx, err
}
