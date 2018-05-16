package main

import (
	"flag"
	"os"
	"log"
)

// Flags
var (
	filePath = flag.String("f", "softwareadvice.json", "Path to the file that needs to be imported")
)

func main() {
	flag.Parse()

	productList, err := readJSON(*filePath)
	if err != nil {
		os.Exit(1)
	}

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


	for _, product := range productList {
		// IF WE HAD A DB SESSION THIS WOULD BE THE CORRECT CALL
		// err = insertRecord(product, mongoDBDialInfo)
		err = insertRecord(product)
		if err != nil {
			log.Printf("Failed inserting record: %s", err)
		}
	}
}
