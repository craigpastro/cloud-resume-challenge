package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("should read a value", func(t *testing.T) {
		resp, err := handler(events.APIGatewayProxyRequest{})
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatal("Non 200 response")
		}

		item := Item{}
		if err := json.Unmarshal([]byte(resp.Body), &item); err != nil {
			t.Fatalf("got error: %v", err)
		}

		if item.Value < 0 {
			t.Fatal("value is less than 0")
		}
	})
}
