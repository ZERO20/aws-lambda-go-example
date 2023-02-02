package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type InputEvent struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

type Response struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

func main() {
	lambda.Start(Handler) // disparador de la lambda
}

func Handler(event InputEvent) (Response, error) {
	fmt.Println("Event: ", event)
	var slink string
	if event.Link == "" {
		slink = "Nada"
	} else {
		slink = event.Link
	}
	return Response{
		Link: slink,
		Key:  event.Key,
	}, nil
}
