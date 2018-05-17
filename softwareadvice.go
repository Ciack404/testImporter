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

// THIS IS HOW THE FUNCTION EXPECTING THE DB SESSION WOULD BE
// func insertRecord(prod Product, collection *mgo.Collection) error {
func (p *Product) insertSoftwareadviceRecord() error {
  report := fmt.Sprintf("importing: Name: \"%s\"; Categories: ", p.Title)
  for i, cat := range p.Categories {
    if i == len(p.Categories)-1 {
      report = fmt.Sprintf("%s%s; ", report, cat)
    } else {
    report = fmt.Sprintf("%s%s ", report, cat)
    }
  }
  report = fmt.Sprintf("%sTwitter: %s", report, p.Twitter)

  fmt.Println(report)

  // err := collection.Insert(p)
	// if err != nil {
	// 	log.Printf("Failed inserting record: %s", err)
	// 	return err
	// }

  fmt.Println("Record successfully inserted")

  return nil
}
