package config

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AWS struct {
		Endpoint string `yaml:"endpoint"`
		Region   string `yaml:"region"`
	} `yaml:"aws"`
}

func LoadLocalConfig() Config {

	file, err := os.Open("environments/local.yaml")
	if err != nil {
		log.Fatalf("Errore nella lettura del file di configurazione: %v", err)
	}
	defer file.Close()

	configData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Errore nella lettura del file di configurazione: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Errore nel parsing del file di configurazione: %v", err)
	}

	return config
}

func InitLocalAwsSession() *dynamodb.DynamoDB {
	// Load configuration
	cfg := LoadLocalConfig()

	// Initialize AWS session
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(cfg.AWS.Endpoint),
		Region:   aws.String(cfg.AWS.Region),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}
	log.Printf("AWS session created")
	// Initialize DynamoDB client
	return dynamodb.New(sess)
}

func InitProdAwsSession() *dynamodb.DynamoDB {
    accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
    secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
    region := os.Getenv("AWS_REGION")

    if accessKey == "" || secretKey == "" || region == "" {
        log.Fatalf("AWS environment variables are not set")
    }

    // Initialize AWS session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if err != nil {
        log.Fatalf("Failed to create AWS session: %v", err)
    }

    log.Println("AWS session for production created successfully")
    return dynamodb.New(sess)
}
