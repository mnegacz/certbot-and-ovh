package main

import (
	"fmt"
	"github.com/ovh/go-ovh/ovh"
	"io/ioutil"
	"log"
	"time"
)

type RecordRequest struct {
	SubDomain string `json:"subDomain"`
	Ttl       int    `json:"ttl"`
	Target    string `json:"target"`
	FieldType string `json:"fieldType"`
}

type Record struct {
	*Record
	Id int `json:"id"`
}

func main() {
	domain := getenv("CERTBOT_DOMAIN")
	log.Printf("domain %s\n", domain)
	validation := getenv("CERTBOT_VALIDATION")
	log.Printf("validation %s\n", validation)

	client := client()
	record := createRecord(client, domain, validation)
	log.Printf("record with id %d for domain %s created", record.Id, domain)
	refreshDomain(client, domain)
	log.Printf("domain %s refreshed", domain)
	persist(domain, validation, record.Id)
	log.Printf("record id for domain %s persisted in file", domain)
	log.Println("sleeping for propagation")
	time.Sleep(30 * time.Second)
}

func persist(domain string, validation string, recordId int) {
	id := fmt.Sprintf("%d", recordId)
	err := ioutil.WriteFile(filePath(domain, validation), []byte(id), 0644)
	assertOk(err)
}

func createRecord(client *ovh.Client, domain string, validation string) Record {
	request := RecordRequest{
		SubDomain: "_acme-challenge",
		Ttl:       120,
		Target:    validation,
		FieldType: "TXT",
	}
	var result Record
	err := client.Post(fmt.Sprintf("/domain/zone/%s/record", domain), request, &result)
	assertOk(err)
	return result
}
