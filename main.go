package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var input string
	fmt.Scan(&input)

	reader, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reader.Name())
}