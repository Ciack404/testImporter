package main

import (
	"flag"
	"os"
	"log"
	"strings"
)

// Flags
var (
	filePath = flag.String("f", "softwareadvice.json", "Path to the file that needs to be imported")
)

func main() {
	flag.Parse()

	// MOCKING DB CREDS AND CONNECTION
	// mongoDBDialInfo := &mgo.DialInfo{
	// 	Addrs:    []string{MongoDBHosts},
	// 	Timeout:  60 * time.Second,
	// 	Database: AuthDatabase,
	// 	Username: AuthUserName,
	// 	Password: AuthPassword,
	// }
	// mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	// if err != nil {
	// 	log.Fatalf("CreateSession: %s\n", err)
	// }


	splittedPath:= strings.Split(*filePath, ".")

	if len(splittedPath) < 2 {
		log.Println("Invalid file: missing extension")
	}
	if splittedPath[len(splittedPath)-1] != "json" && splittedPath[len(splittedPath)-1] != "yaml" {
			log.Println("Unexpected file extension (expecting .json or .yaml)")
	}

	if splittedPath[len(splittedPath)-1] == "json" {
		productList := Softwareadvice{}
		err := productList.readJSON(*filePath)
		if err != nil {
			os.Exit(1)
		}

		for _, product := range productList.Products {
			// IF WE HAD A DB SESSION COLLECTION THIS WOULD BE THE CORRECT CALL
			// collection := mongoSession.DB("TestDatabase").C("softwareadvice")
			// err = insertRecord(product, collection)
			err = insertSoftwareadviceRecord(product)
			if err != nil {
				log.Printf("Failed inserting record: %s", err)
			}
		}
	} else {
		productList := Capterra{}
		err := productList.readYAML(*filePath)
		if err != nil {
			os.Exit(1)
		}

		for _, product := range productList {
			// IF WE HAD A DB SESSION COLLECTION THIS WOULD BE THE CORRECT CALL
			// collection := mongoSession.DB("TestDatabase").C("capterra")
			// err = insertRecord(product, collection)
			err = insertCapterraRecord(product)
			if err != nil {
				log.Printf("Failed inserting record: %s", err)
			}
		}
	}
}
