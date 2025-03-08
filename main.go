package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"task-service/config" // Update this import path to match your project structure
)

type Tasks struct {
	TaskID       string `json:"task_id"`
	Description  string `json:"description"`
	StartingTime string `json:"starting_time"`
	Duration     int    `json:"duration"`
	Status       string `json:"status"`
}

func main() {
	// Carica il file di configurazione
	config := config.LoadConfig()

	// Crea una nuova sessione AWS per DynamoDB
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(config.AWS.Endpoint),
		Region:   aws.String(config.AWS.Region),
	})
	if err != nil {
		log.Fatalf("Errore nella creazione della sessione AWS: %v", err)
	}

	// Crea un client DynamoDB
	svc := dynamodb.New(sess)

	// Inserisci un task in DynamoDB
	task := Tasks{
		TaskID:       "task_001",
		Description:  "Complete project documentation",
		StartingTime: "08:00",
		Duration:     120,
		Status:       "PENDING",
	}

	// Converte la struttura in un formato compatibile con DynamoDB
	av, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		log.Fatalf("Errore nella conversione della struttura in attributi DynamoDB: %v", err)
	}

	// Inserisci l'elemento nella tabella "Tasks"
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("Tasks"),
		Item:      av,
	})
	if err != nil {
		log.Fatalf("Errore nell'inserimento del task in DynamoDB: %v", err)
	}

	fmt.Println("Task inserito con successo in DynamoDB!")
}
