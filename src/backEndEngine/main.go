package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	fmt.Println("Starting backend Engine......")

	router := mux.NewRouter()

	// define the handler Functions
	router.HandleFunc("/getEntries/", getEntries)
	router.HandleFunc("/addEntry/", addEntry)
	router.HandleFunc("/removeEntry/", removeEntry)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"content_type", "Accept", "Accept-Language", "Content-Type"},
	}).Handler(router)

	// start listening
	log.Fatal(http.ListenAndServe(":8080", handler))
}
