package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Initializing file name variable
	var fName string
	fmt.Println("Enter the file name: ")
	fmt.Scan(&fName)

	// Read file if it exists
	f,err := os.ReadFile(fName)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(len(f))

	// MAP TO KEEP A COUNT OF THE CHARACTERS
	charCount := make(map[string]int64)

	for _,c := range f{
		charCount[string(c)]++
	}

	fmt.Println(charCount)

}