package main

import (
	"encoding/json"
	"io/ioutil"
)

const credLocation = "./credentials.json"

type Creds struct {
	League      string `json:"league"`
	POESESSID   string `json:"POESESSID"`
	AccountName string `json:"accountName"`
}

func ReadCreds() *Creds {
	return ReadJson(credLocation, &Creds{})
}

func ReadJson[T any](file string, item *T) *T {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContents, item)

	if err != nil {
		panic(err)
	}

	return item
}
