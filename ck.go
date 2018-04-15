package main

import (
	"fmt"
	"github.com/ovh/go-ovh/ovh"
)

func main() {
	// Create a client using credentials from config files or environment variables
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}
	ckReq := client.NewCkRequest()

	ckReq.AddRule("POST", "/domain/zone/*/record")
	ckReq.AddRule("DELETE", "/domain/zone/*/record/*")
	ckReq.AddRule("POST", "/domain/zone/*/refresh")

	// Run the request
	response, err := ckReq.Do()
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	// Print the validation URL and the Consumer key
	fmt.Printf("Generated consumer key: %s\n", response.ConsumerKey)
	fmt.Printf("Please visit %s to validate it\n", response.ValidationURL)
}
