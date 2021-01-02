package main

/*
	Function: main
	Author	: Glen Small
	Date	: 16th Decemer 2020

	Params	: none
	Returns	: nil

	Purpose : Main entrypoint for the application
*/
func main() {

	//TODO: This all needs to be moved but not sure on the final location yet.

	/*
		displayBanner()

		godotenv.Load()
		consoleout.printSuccess("Loading OS Environment Handler")

		router := mux.NewRouter()
		consoleout.printSuccess("Creating MUX Router")

		// define the handler Functions
		router.HandleFunc("/getEntries/", getEntries)
		router.HandleFunc("/addEntry/", addEntry)
		router.HandleFunc("/removeEntry/", removeEntry)
		consoleout.printSuccess("Adding routes to MUX router")

		handler := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PATCH", "OPTIONS"},
			AllowedHeaders: []string{"content_type", "Accept", "Accept-Language", "Content-Type"},
		}).Handler(router)
		consoleout.printSuccess("Adding CORS default headers to the router")

		consoleout.printInfo("Attempting to connect to the backend Database")
		myClient, myContext, err := initDB()

		if err != nil {
			consoleout.printError("Exiting ......")
			return
		}

		defer func() {
			if err = myClient.Disconnect(*myContext); err != nil {
				consoleout.printError(fmt.Sprintf("Error disconnecting from DB - %s", err))
			}
		}()

		printInfo("Starting HTTP Listener on [" + os.Getenv("ENGINE_LISTENER") + "]")
		// start listening
		log.Fatal(http.ListenAndServe(os.Getenv("ENGINE_LISTENER"), handler))
	*/
}
