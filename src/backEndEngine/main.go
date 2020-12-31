package main

import (
	"fmt"
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

	fmt.Println("Loading OS Environment handler ......")
	godotenv.Load()

	fmt.Println("Creating MUX Router .......")
	router := mux.NewRouter()

	fmt.Println("Adding router handlers to the router .......")
	// define the handler Functions
	router.HandleFunc("/getEntries/", getEntries)
	router.HandleFunc("/addEntry/", addEntry)
	router.HandleFunc("/removeEntry/", removeEntry)

	fmt.Println("Adding CORS default headers to the router ..........")
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"content_type", "Accept", "Accept-Language", "Content-Type"},
	}).Handler(router)

	fmt.Println("Starting HTTP Listener on [", os.Getenv("ENGINE_LISTENER"), "]")
	// start listening
	log.Fatal(http.ListenAndServe(os.Getenv("ENGINE_LISTENER"), handler))
}
