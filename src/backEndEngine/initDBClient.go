package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
)

// function to open up a mongo conection
func initDB() (*mongo.Client, *context.Context, error) {

	// mongodb://myDBReader:D1fficultP%40ssw0rd@mongodb0.example.com:27017/?authSource=admin
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_DB"))

	printInfo(fmt.Sprintf("Using Connection String [%s]", connectionString))
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))

	if err != nil {
		printError(fmt.Sprintf("Failed to create the Mongo Client - %s", err))
		log.Fatal(err)
		return nil, nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = nil
	err = client.Connect(ctx)

	if err != nil {
		printError(fmt.Sprintf("Failed to connect - %s", err))
		log.Fatal(err)
		return nil, nil, err
	}

	return client, &ctx, err
}
