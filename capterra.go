package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
)

type Capterra []CapterraProduct

type CapterraProduct struct {
  Tags string `yaml:"tags"`
  Name string `yaml:"name"`
  Twitter string `yaml:"twitter"`
}

func (c *Capterra) unmarshalYAML(file []byte) error {
    err := yaml.Unmarshal(file, c)
    if err != nil {
        fmt.Printf("Unmarshal: %v", err)
        return err
    }

    return nil
}

// THIS IS HOW THE FUNCTION EXPECTING THE DB SESSION WOULD BE
// func (p *CapterraProduct) insertRecord(collection *mgo.Collection) error {
func (p *CapterraProduct) insertCapterraRecord() error {
  fmt.Printf("importing: Name: \"%s\"; Categories: %s; Twitter: %s\n", p.Name, p.Tags, p.Twitter)

  // err := collection.Insert(p)
	// if err != nil {
	// 	log.Printf("Failed inserting record: %s", err)
	// 	return err
	// }

  fmt.Println("Record successfully inserted")

  return nil
}
