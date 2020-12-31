package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	displayBanner()

	godotenv.Load()
	printSuccess("Loading OS Environment Handler")

	router := mux.NewRouter()
	printSuccess("Creating MUX Router")

	// define the handler Functions
	router.HandleFunc("/getEntries/", getEntries)
	router.HandleFunc("/addEntry/", addEntry)
	router.HandleFunc("/removeEntry/", removeEntry)
	printSuccess("Adding routes to MUX router")

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"content_type", "Accept", "Accept-Language", "Content-Type"},
	}).Handler(router)
	printSuccess("Adding CORS default headers to the router")

	printInfo("Starting HTTP Listener on [" + os.Getenv("ENGINE_LISTENER") + "]")
	// start listening
	log.Fatal(http.ListenAndServe(os.Getenv("ENGINE_LISTENER"), handler))
}
