package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:8080"

func main() {
	lc := loginClient(url+"/v1/login", "contacto@ed.team", "123456")
	fmt.Println(lc)

	person := Person{
		Name:        "Andrey",
		Age:         21,
		Communities: []Community{Community{Name: "EDteam"}},
	}
	gr := createPerson(url+"/v1/persons", lc.Data.Token, &person)
	fmt.Println(gr)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Rquest: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Response: %v", err)
	}

	return response
}
