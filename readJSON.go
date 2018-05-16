package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Softwareadvice represent the main structure expected as imput
type Softwareadvice struct {
	Products []Product `json:"products"`
}

// Product represent a product in the feed provided
type Product struct {
	Categories []string `json:"categories"`
	Twitter    string   `json:"twitter"`
	Title      string   `json:"title"`
}

func (s *Softwareadvice) readJSON(filePath string) error {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	json.Unmarshal(raw, &s)
	return nil
}
