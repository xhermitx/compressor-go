package main

import (
	"fmt"
	"log"
	"os"

	huffman "example.com/Compressor/Huffman"
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
	charCount := make(map[rune]int)

	for _,c := range f{
		charCount[rune(c)]++
	}

	fmt.Println(charCount)

	sampleTree := huffman.BuildTree(charCount)

	huffman.PrintCodes(sampleTree, []byte{})

}