package main

import (
	"fmt"
	"github.com/ovh/go-ovh/ovh"
	"log"
	"os"
)

func refreshDomain(client *ovh.Client, domain string) {
	err := client.Post(fmt.Sprintf("/domain/zone/%s/refresh", domain), struct{}{}, &struct{}{})
	assertOk(err)
}

func client() *ovh.Client {
	client, err := ovh.NewDefaultClient()
	assertOk(err)
	return client
}

func getenv(name string) string {
	value := os.Getenv(name)
	assertNotEmpty(name, value)
	return value
}

func assertOk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func assertNotEmpty(name string, value string) {
	if len(value) == 0 {
		log.Fatalf("enviroment variable %s not found\n", name)
	}
}

func filePath(domain string, validation string) string {
	return "/tmp/" + domain + "_" + validation
}
