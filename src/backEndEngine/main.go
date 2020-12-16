package main

import (
	"fmt"
	"net/http"
)

/*
	Function: main
	Author	: Glen Small
	Date	: 16th Decemer 2020

	Params	: none
	Returns	: nil

	Purpose : Main entrypoint for the application
*/
func main() {

	fmt.Println("Starting webServer......")

	// define the handler Functions
	http.HandleFunc("/getEntries/", getEntries)
	http.HandleFunc("/addEntry/", addEntry)
	http.HandleFunc("/removeEntry/", removeEntry)


	// start listening
	http.ListenAndServe(":8080", nil)
}
