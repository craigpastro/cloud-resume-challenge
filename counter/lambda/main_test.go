package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("increment should return 200", func(t *testing.T) {
		resp, err := handler(events.APIGatewayProxyRequest{})
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatal("Non 200 response")
		}
	})

	t.Run("increment increments", func(t *testing.T) {
		resp, _ := handler(events.APIGatewayProxyRequest{})

		item := Item{}
		json.Unmarshal([]byte(resp.Body), &item)
		current := item.Value

		resp, _ = handler(events.APIGatewayProxyRequest{})
		json.Unmarshal([]byte(resp.Body), &item)
		next := item.Value

		if next != current+1 {
			t.Fatalf("failed to increment counter. current=%d, next=%d", current, next)
		}
	})
}
