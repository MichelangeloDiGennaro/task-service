package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Task struct {
	TaskID       string `json:"task_id" dynamodbav:"task_id"`
	Description  string `json:"description"`
	StartingTime string `json:"starting_time"`
	Duration     int    `json:"duration"`
	Status       string `json:"status"`
}

func CreateTaskHandler(db *dynamodb.DynamoDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// check if the Task struct is valid by converting it into a format that is compatible
		// with the DynamoDB: ensuring that the data types and attribute names match those defined in the DynamoDB table.
		av, err := dynamodbattribute.MarshalMap(task)
		if err != nil {
			http.Error(w, "Failed to marshal task", http.StatusInternalServerError)
			return
		}

		_, err = db.PutItem(&dynamodb.PutItemInput{
			TableName:           aws.String("Tasks"),
			Item:                av,
			ConditionExpression: aws.String("attribute_not_exists(task_id)"), // Ensure no duplicate task_id
		})
		if err != nil {
			http.Error(w, "Failed to create task: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}
