package main

import (
  "fmt"
)

// THIS IS HOW THE FUNCTION EXPECTING THE DB SESSION WOULD BE
// func insertRecord(prod Product, collection *mgo.Collection) error {
func insertSoftwareadviceRecord(prod Product) error {
  report := fmt.Sprintf("importing: Name: \"%s\"; Categories: ", prod.Title)
  for i, cat := range prod.Categories {
    if i == len(prod.Categories)-1 {
      report = fmt.Sprintf("%s%s; ", report, cat)
    } else {
    report = fmt.Sprintf("%s%s ", report, cat)
    }
  }
  report = fmt.Sprintf("%sTwitter: %s", report, prod.Twitter)

  fmt.Println(report)

  // err := collection.Insert(prod)
	// if err != nil {
	// 	log.Printf("Failed inserting record: %s", err)
	// 	return err
	// }

  fmt.Println("Record successfully inserted")

  return nil
}

// THIS IS HOW THE FUNCTION EXPECTING THE DB SESSION WOULD BE
// func insertRecord(prod Capterra, collection *mgo.Collection) error {
func insertCapterraRecord(prod CapterraProduct) error {
  fmt.Printf("importing: Name: \"%s\"; Categories: %s; Twitter: %s", prod.Name, prod.Tags, prod.Twitter)

  // err := collection.Insert(prod)
	// if err != nil {
	// 	log.Printf("Failed inserting record: %s", err)
	// 	return err
	// }

  fmt.Println("Record successfully inserted")

  return nil
}
