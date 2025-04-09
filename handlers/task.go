package handlers

import (
	"context"
	"fmt"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"task-service/config"
)

type Task struct {
	TaskID       string `json:"task_id" dynamodbav:"task_id"`
	Description  string `json:"description"`
	StartingTime string `json:"starting_time"`
	Duration     int    `json:"duration"`
	Status       string `json:"status"`
}

func CreateTaskHandler(dctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	var task Task
	var db *dynamodb.DynamoDB
	db = config.InitProdAwsSession()
	err := json.Unmarshal([]byte(request.Body), &task)

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request payload",
		}, nil
	}

	// check if the Task struct is valid by converting it into a format that is compatible
	// with the DynamoDB: ensuring that the data types and attribute names match those defined in the DynamoDB table.
	av, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf(`Failed to marshal task`),
			}, nil
	}

	_, err = db.PutItem(&dynamodb.PutItemInput{
		TableName:           aws.String("Tasks"),
		Item:                av,
		ConditionExpression: aws.String("attribute_not_exists(task_id)"), // Ensure no duplicate task_id
	})
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf(`Failed to create task`+err.Error()),
			}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf(`task created successfully`),
		}, nil
}

func UnhandledMethod(dctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       fmt.Sprintf(`UnhandledMethod`),
		}, nil
}
