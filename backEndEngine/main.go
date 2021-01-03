package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"consolewriter"
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
	consolewriter.PrintSuccess("Loading OS Environment Handler")

	router := mux.NewRouter()
	consolewriter.PrintSuccess("Creating MUX Router")

	// define the handler Functions
	router.HandleFunc("/getEntries/", getEntries)
	router.HandleFunc("/addEntry/", addEntry)
	router.HandleFunc("/removeEntry/", removeEntry)
	consolewriter.PrintSuccess("Adding routes to MUX router")

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"content_type", "Accept", "Accept-Language", "Content-Type"},
	}).Handler(router)
	consolewriter.PrintSuccess("Adding CORS default headers to the router")

	consolewriter.PrintInfo("Attempting to connect to the backend Database")
	//TODO:  Need to add the db connection stuff here

	/*
		if err != nil {
			consolewriter.PrintError("Exiting ......")
			return
		}

		defer func() {
			if err = myClient.Disconnect(*myContext); err != nil {
				consolewriter.PrintError(fmt.Sprintf("Error disconnecting from DB - %s", err))
			}
		}()
	*/

	consolewriter.PrintInfo("Starting HTTP Listener on [" + os.Getenv("ENGINE_LISTENER") + "]")
	// start listening
	log.Fatal(http.ListenAndServe(os.Getenv("ENGINE_LISTENER"), handler))
}
