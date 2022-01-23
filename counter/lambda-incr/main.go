package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/kelseyhightower/envconfig"
)

const (
	counter   = "Counter"
	tableName = "cloud-resume-challenge-counter"
	value     = "Value"
)

var ErrInternalServer = events.APIGatewayProxyResponse{StatusCode: 500}

type Config struct {
	Region string `default:"us-west-2"`
	Local  bool   `default:"false"`
}

type Item struct {
	Value int `json:"value"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	client, err := newDDBClient()
	if err != nil {
		return ErrInternalServer, err
	}

	update := expression.Add(expression.Name(value), expression.Value(1))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return ErrInternalServer, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(counter),
			},
		},
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              aws.String("UPDATED_NEW"),
	}
	resp, err := client.UpdateItem(input)
	if err != nil {
		return ErrInternalServer, err
	}

	var item Item
	if err := dynamodbattribute.UnmarshalMap(resp.Attributes, &item); err != nil {
		return ErrInternalServer, err
	}

	body, err := json.Marshal(item)
	if err != nil {
		return ErrInternalServer, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func newDDBClient() (*dynamodb.DynamoDB, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal("error reading config", err)
	}

	c := aws.Config{Region: aws.String(config.Region)}
	if config.Local {
		c.Endpoint = aws.String("http://localhost:8000")
	}
	sess, err := session.NewSession(&c)
	if err != nil {
		return nil, err
	}

	return dynamodb.New(sess), nil
}

func main() {
	lambda.Start(handler)
}
