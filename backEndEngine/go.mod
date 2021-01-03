module backendEngine

go 1.15

require (
	DBLayer v0.0.0 // indirect
	consolewriter v0.0.0 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	github.com/rs/cors v1.7.0
	go.mongodb.org/mongo-driver v1.4.4 // indirect
)

replace consolewriter => ./consolewriter

replace DBLayer => ./DBLayer