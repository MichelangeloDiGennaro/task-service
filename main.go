package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"task-service/config"   // Update this import path to match your project structure
	"task-service/handlers" // Import your handlers package

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
)

func main() {

	// Define a command-line flag for the environment
	env := flag.String("env", "local", "Environment to run the application in (local or prod)")
	flag.Parse()

	// Print the environment variable
	fmt.Printf("Environment: %s\n", *env)

	var db *dynamodb.DynamoDB
	//env := os.Getenv("APP_ENV")

	if *env == "local" {
		fmt.Printf("local environment")

		db = config.InitLocalAwsSession()
	} else {
		fmt.Printf("Running in prod environmen")
		db = config.InitProdAwsSession()
	}

	// Initialize router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/add-task", handlers.CreateTaskHandler(db)).Methods("POST")
	//router.HandleFunc("/tasks/{id}", handlers.GetTaskHandler(db)).Methods("GET")
	// Add more routes as needed

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
