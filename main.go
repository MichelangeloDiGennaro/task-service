package main

import (
	"flag"
	"fmt"
	//"log"
	//"net/http"
	//"os"
	"context"
	"task-service/handlers"
	"github.com/aws/aws-lambda-go/lambda"
	//"github.com/gorilla/mux"
	"github.com/aws/aws-lambda-go/events"
)

func main() {

	// Define a command-line flag for the environment
	env := flag.String("env", "local", "Environment to run the application in (local or prod)")
	flag.Parse()

	// Print the environment variable
	fmt.Printf("Environment: %s\n", *env)

	//env := os.Getenv("APP_ENV")


	//router.HandleFunc("/tasks/{id}", handlers.GetTaskHandler(db)).Methods("GET")
	// Add more routes as needed

	// Start the server
	if *env == "local" {
		//fmt.Printf("local environment")
		// Initialize router
		//router := mux.NewRouter()

		// Register routes
		//router.HandleFunc("/add-task", handlers.CreateTaskHandler).Methods("POST")
		//db = config.InitLocalAwsSession()
		//port := os.Getenv("PORT")
		//if port == "" {
			//port = "8080"
		//}

		//fmt.Printf("Server is running on port %s\n", port)
		//log.Fatal(http.ListenAndServe(":"+port, router))

	} else {
		fmt.Printf("Running in prod environmen")
		// Start the Lambda handler
		lambda.Start(lambdaHandler)
	}

}

func lambdaHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
		case "POST":
			return handlers.CreateTaskHandler(context.Background(), req)
		default:
			return handlers.UnhandledMethod(context.Background(), req)
	}
}
