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

func readJSON(filePath string) ([]Product, error) {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return []Product{}, err
	}

	var c Softwareadvice
	json.Unmarshal(raw, &c)
	return c.Products, nil
}
