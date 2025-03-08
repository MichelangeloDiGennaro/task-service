package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"

	"task-service/config"   // Update this import path to match your project structure
	"task-service/handlers" // Import your handlers package
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize AWS session
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(cfg.AWS.Endpoint),
		Region:   aws.String(cfg.AWS.Region),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Initialize DynamoDB client
	db := dynamodb.New(sess)

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
