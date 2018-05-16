package main

import (
  "fmt"
)

// THIS IS HOW THE FUNCTION EXPECTING THE DB SESSION WOULD BE
// func insertRecord(Product, *mgo.Session) error {
func insertRecord(prod Product) error {
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

  return nil
}
