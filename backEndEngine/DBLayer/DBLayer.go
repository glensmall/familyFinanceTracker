package dblayer

import (
	"fmt"
	"os"
	"time"

	"backendEngine / consolewriter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
)

// a couple of global variables I'll be using
var client *mongo.Client
var ctx context.Context

/*
	This is a structure that defines what each document in the DB will look like
	Will use this when we read and write to mongo
*/
type Transaction Struct {
	ID int64 `json:"id:omitempty"`,
	Description string `json:"description:omitempty`,
	Date string `json:"date:omitempty`,
	Amount float32 `json:"amount:omitempty`
} 

/*
	Function: Connect
	Author	: Glen Small
	Date	: 3rd January 2021

	Params	: none
	Returns	: error

	Purpose : Creates the requierd objects and connects to the defines database
			  The connection string varibales will be read from .env and formatted as needed
*/
func Connect() error {

	// mongodb://myDBReader:D1fficultP%40ssw0rd@mongodb0.example.com:27017/?authSource=admin
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_DB"))

	consolewriter.PrintInfo(fmt.Sprintf("Using Connection String [%s]", connectionString))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	// since Connect() is non-blocking at this stage, lets test to see if we actually got that last bit right
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		consolewriter.PrintError(fmt.Sprintf("Error Initialising the DB connection - %s", err))
		return err
	}

	consolewriter.PrintSuccess("Database found and connection established")

	return err
}


/*
	Function: Disconnect
	Author	: Glen Small
	Date	: 3rd January 2021

	Params	: none
	Returns	: error

	Purpose : disconnects from the DB and closes the connection
*/
func Disconnect() (error){

	err := client.Disconnect(ctx)

	if err != nil {
		consolewriter.PrintError(fmt.sprintf("Error disconnecting from the database - %s", err))
		return err
	} 

	consolewriter.PrintSuccess("Database Disconnected")

	return err
}