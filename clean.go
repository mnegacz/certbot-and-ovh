package main

import (
	"fmt"
	"github.com/ovh/go-ovh/ovh"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	domain := getenv("CERTBOT_DOMAIN")
	log.Printf("domain %s\n", domain)
	validation := getenv("CERTBOT_VALIDATION")
	log.Printf("validation %s\n", validation)

	filePath := filePath(domain, validation)
	id := readId(filePath)

	client := client()
	deleteRecord(client, domain, id)
	log.Printf("record with id %d for domain %s deleted", id, domain)
	refreshDomain(client, domain)
	log.Printf("domain %s refreshed", domain)
	remove(filePath)
	log.Printf("file with record id for domain %s removed", domain)
}

func readId(filePath string) int {
	assertExists(filePath)
	dat, err := ioutil.ReadFile(filePath)
	assertOk(err)
	id, err := strconv.Atoi(string(dat))
	assertOk(err)
	return id
}
func remove(filePath string) {
	err := os.Remove(filePath)
	assertOk(err)
}
func deleteRecord(client *ovh.Client, domain string, id int) {
	err := client.Delete(fmt.Sprintf("/domain/zone/%s/record/%d", domain, id), struct{}{})
	assertOk(err)
}

func assertExists(filePath string) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
}
