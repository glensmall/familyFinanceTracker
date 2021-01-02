package main

import (
	"context"
	"net/http"
	"time"
)

/*
	Function: getEntries
	Author	: Glen Small
	Date	: 16th December 2020

	Params	: http.ResponseWriter - for communication back to the client
			  http.Request - pointer to the request object that has been made

	Returns : returns a list of all entries in the database

	Purpose : Backend handler to remove an entry from the database
*/
func getEntries(w http.ResponseWriter, r *http.Request) {

	// establish a context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

}
