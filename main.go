package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Tasks struct {
	TaskID       string `json:"task_id"`
	Description  string `json:"description"`
	StartingTime string `json:"starting_time"`
	Duration     int    `json:"duration"`
	Status       string `json:"status"`
}

func main() {
	// Crea una nuova sessione AWS per DynamoDB Local
	sess, err := session.NewSession(&aws.Config{
		// Impostiamo l'endpoint per DynamoDB Local
		Endpoint: aws.String("http://localhost:8000"), // Imposta l'endpoint di DynamoDB Local
		Region:   aws.String("us-west-1"),             // Puoi usare qualsiasi regione, non ha importanza per DynamoDB Local
	})
	if err != nil {
		log.Fatalf("Errore nella creazione della sessione AWS: %v", err)
	}

	// Crea un client DynamoDB
	svc := dynamodb.New(sess)

	// Inserisci un task in DynamoDB Local
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

	fmt.Println("Task inserito con successo in DynamoDB Local!")
}
