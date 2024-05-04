package main

import (
	"fmt"
	"os"

	encoder "example.com/Compressor/Encoder"
	fileIO "example.com/Compressor/FileIO"
	huffman "example.com/Compressor/Huffman"
)

func main() {
	//Initializing file name variable
	args := os.Args[1:]
	
	FLAG := args[0]
	inputFile := args[1]
	outpuFile := args[2]

	switch FLAG{
	case "-e":
		content := fileIO.ReadFromFile(inputFile)

		// MAP TO KEEP A COUNT OF THE CHARACTERS
		charCount := make(map[rune]int)

		for _,c := range content{
			charCount[rune(c)]++
		}

		// BUILD THE HUFFMAN TREE FOR ENCODING THE CHARACTERS
		sampleTree := huffman.BuildTree(charCount)

		// CREATE AN ENCODER MAP OF CHARACTER:BIT_STRING
		encoderMap := huffman.GenerateCodes(sampleTree, []byte{}, make(map[rune]string))

		// WRITE THE HEADER TO THE FILE
		fileIO.WriteHeader(outpuFile, encoderMap)

		// ENCODE THE FILE CONTENTS
		encoder.Encode(inputFile, outpuFile, encoderMap)
	
	case "-o":

		// READ HEADER FROM THE FILE
		encoderMap := fileIO.ReadHeader(inputFile)

		encoder.Decode(inputFile, outpuFile, encoderMap)
	
	default:
		fmt.Println("Please provide correct arguments")
	}
}