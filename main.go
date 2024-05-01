package main

import (
	"fmt"
	"log"
	"os"

	fileIO "example.com/Compressor/FileIO"
	huffman "example.com/Compressor/Huffman"
)

func main() {
	//Initializing file name variable
	var fileName string
	fmt.Println("Enter the file name: ")
	fmt.Scan(&fileName)

	// Read file if it exists
	f,err := os.ReadFile(fileName)
	if err!=nil{
		log.Fatal(err)
	}

	//Initializing file name variable
	var compressedFileName string
	fmt.Println("Enter the file name: ")
	fmt.Scan(&compressedFileName)

	// MAP TO KEEP A COUNT OF THE CHARACTERS
	charCount := make(map[rune]int)

	for _,c := range f{
		charCount[rune(c)]++
	}

	// BUILD THE HUFFMAN TREE FOR ENCODING THE CHARACTERS
	sampleTree := huffman.BuildTree(charCount)

	// CREATE AN ENCODER MAP OF CHARACTER:BIT_STRING
	encoderMap := huffman.GenerateCodes(sampleTree, []byte{}, make(map[rune]string))

	// ENCODE THE FILE CONTENTS
	fileIO.Encoder(f, encoderMap, charCount, compressedFileName)
	
	// COMPARING THE FILE SIZE BEFORE AND AFTER ENCODING
	fmt.Println("\n\nOriginal File length: ",len(f))
	// fmt.Println("Byte encoded File length: ",len(byteArray))

	// // WRITE THE BYTE ARRAY TO FILE
	// writeToFile(byteArray)
}