package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	// a real simple routine to read the content of our test json file
	fileContents, err := ioutil.ReadFile(".\\testData.json")

	if err != nil {
		fmt.Println("Error reading the testJson File", err)
	}

	// set the headers properly
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application.json")

	// redner the file contents as a string back to the caller
	fmt.Fprintln(w, string(fileContents))

	// some console output.
	fmt.Println("Serving Json")
}
