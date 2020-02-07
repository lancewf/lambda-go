package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Automate struct {
	URL      string
	APIToken string
}

// HandleRequest - handles requests
func HandleRequest(ctx context.Context, message json.RawMessage) (string, error) {

	for _, automate := range getAutomates() {
		req, err := http.NewRequest("POST", automate.URL, bytes.NewReader(message))
		if err != nil {
			fmt.Printf("error sending request %v", err)
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("api-token", automate.APIToken)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("error sending request %v", err)
			continue
		}

		if resp.StatusCode != 200 {
			fmt.Printf("non 200 status code %d", resp.StatusCode)
		}
	}

	return "chef demo", nil
}

func main() {
	lambda.Start(HandleRequest)
}

// Pull from a data store
func getAutomates() []Automate {
	return []Automate{
		{
			URL:      "http://requestbin.sjcmmsn.com/123e2vh1",
			APIToken: "fake-token-1",
		},
		{
			URL:      "http://requestbin.sjcmmsn.com/1o68e1k1",
			APIToken: "fake-token-2",
		},
	}
}
