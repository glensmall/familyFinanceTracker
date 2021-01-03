module backEndEngine

go 1.15

require (
	consolewriter v0.0.0
	dblayer v0.0.0
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	github.com/rs/cors v1.7.0
)

replace consolewriter => ./consolewriter

replace dblayer => ./dblayer
