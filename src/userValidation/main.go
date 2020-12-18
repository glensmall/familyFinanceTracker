package main

import (
	"fmt"
	"net/http"
)

/*
	Function: main
	Author	: Glen Small
	Date	: 18th Decemer 2020

	Params	: none
	Returns	: nil

	Purpose : Main entrypoint for the application
*/
func main() {

	fmt.Println("Starting user Validation Service......")

	// define the handler Functions
	http.HandleFunc("/validateToken/", validateToken)

	// start listening
	http.ListenAndServe(":9000", nil)
}
