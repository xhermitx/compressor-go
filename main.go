package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	huffman "example.com/Compressor/Huffman"
)

// RETURNS AN ENCODED STRING OF BITS USING THE ENCODER MAP
func encoder(file_content []byte, encoder_map map[rune]string)string{
	encoded_bits := ""

	for _,char := range file_content{
		encoded_bits += encoder_map[rune(char)]
	}

	return encoded_bits
}

// CONVERTS THE ENCODED STRING OF BITS INTO A BYTE ARRAY TO WRITE TO A FILE
func bit_to_byte(bit_string string) []byte{
	
	var res []byte
	
	if len(bit_string)%8 !=0{
		padding := 8-(len(bit_string)%8)
		for i:=0;i<padding;i++{
			bit_string += "0"
		}
	}

	for i:=0;i<len(bit_string);i+=8{
		b,_ := strconv.ParseUint(bit_string[i:i+8], 2, 8)
		res = append(res, byte(b))
	}

	return res
}

func main() {
	//Initializing file name variable
	var file_name string
	fmt.Println("Enter the file name: ")
	fmt.Scan(&file_name)

	// Read file if it exists
	f,err := os.ReadFile(file_name)
	if err!=nil{
		log.Fatal(err)
	}

	// MAP TO KEEP A COUNT OF THE CHARACTERS
	char_count := make(map[rune]int)

	for _,c := range f{
		char_count[rune(c)]++
	}

	// BUILD THE HUFFMAN TREE FOR ENCODING THE CHARACTERS
	sample_tree := huffman.BuildTree(char_count)

	// CREATE AN ENCODER MAP OF CHARACTER:BIT_STRING
	encoder_map := huffman.GenerateCodes(sample_tree, []byte{}, make(map[rune]string))

	// ENCODE THE FILE CONTENTS
	encoded_bits := encoder(f, encoder_map)
	byte_array := bit_to_byte(encoded_bits)

	// COMPARING THE FILE SIZE BEFORE AND AFTER ENCODING
	fmt.Println("\n\nOriginal File length: ",len(f))
	fmt.Println("Byte encoded File length: ",len(byte_array))
}