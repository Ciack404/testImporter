package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"io/ioutil"
)

type Capterra []CapterraProduct

type CapterraProduct struct {
  Tags string `yaml:"tags"`
  Name string `yaml:"name"`
  Twitter string `yaml:"twitter"`
}

func (c *Capterra) readYAML(filePath string) error {
    yamlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Failed reading yaml file: %s ", err)
        return err
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        fmt.Printf("Unmarshal: %v", err)
        return err
    }

    return nil
}

// THIS IS HOW THE FUNCTION EXPECTING THE DB SESSION WOULD BE
// func (p *CapterraProduct) insertRecord(collection *mgo.Collection) error {
func (p *CapterraProduct) insertCapterraRecord error {
  fmt.Printf("importing: Name: \"%s\"; Categories: %s; Twitter: %s", p.Name, p.Tags, p.Twitter)

  // err := collection.Insert(p)
	// if err != nil {
	// 	log.Printf("Failed inserting record: %s", err)
	// 	return err
	// }

  fmt.Println("Record successfully inserted")

  return nil
}
