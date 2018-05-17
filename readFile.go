package main

import (
  "fmt"
  "io/ioutil"
)

func readFile(filePath string) ([]byte, error) {
  sourceFile, err := ioutil.ReadFile(filePath)
  if err != nil {
      fmt.Printf("Failed reading yaml file: %s ", err)
      return []byte{}, err
  }
  return sourceFile, nil
}
