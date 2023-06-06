package utils

import (
	"fmt"

	"gopkg.in/resty.v1"
)

// TODO: make it scalable
const baseUrlMachineLearning = "https://ml.arvigo.site"

func FetchMachineLearningAPI(method, path string, body interface{}) (res []byte, err error) {
	// Create a new Resty client
	client := resty.New()
	url := fmt.Sprintf("%s/%s", baseUrlMachineLearning, path)

	// Create the request object
	restyReq := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-KEY", "05d27624-e862-4127-afc8-d382a137ec52")

	// Set the request method
	var response *resty.Response
	switch method {
	case "GET":
		response, err = restyReq.Get(url)
	case "POST":
		response, err = restyReq.SetBody(body).Post(url)
	case "PUT":
		response, err = restyReq.SetBody(body).Put(url)
	case "DELETE":
		response, err = restyReq.Delete(url)
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	// Check for any errors
	if err != nil {
		return nil, fmt.Errorf("error occurred during the request: %v", err)
	}

	// Check the response status code
	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode())
	}

	// Return the response body
	return response.Body(), nil
}
