package encoder

import (
	"fmt"
	"strconv"
	"strings"

	fileIO "example.com/Compressor/FileIO"
)

const CHUNK_SIZE = 20000

func Encode(inputFile string, outputFile string,  encoderMap map[rune]string){

	encodedBits := ""
	extra := 0
	carryBits := ""

	fileContent := fileIO.ReadFromFile(inputFile)

	for _, b := range fileContent {
		if extra > 0 {
			encodedBits += carryBits
			carryBits = ""
			extra = 0
		}
		// fmt.Println(encoderMap[rune(b)])
		encodedBits += encoderMap[rune(b)]

		if len(encodedBits) >= CHUNK_SIZE {
			extra = len(encodedBits) % CHUNK_SIZE
			carryBits = encodedBits[len(encodedBits)-extra:]
			encodedBits = encodedBits[:len(encodedBits)-extra]
			bitToByte(encodedBits, outputFile)
			encodedBits = ""
		}
	}
	bitToByte(encodedBits, outputFile)
}

// CONVERTS THE ENCODED STRING OF BITS INTO A BYTE ARRAY TO WRITE TO A FILE
func bitToByte(bitString string, outputFile string) {

	var res []byte
	// CALCULATING THE PADDING BITS
	if len(bitString)%8 != 0 {
		padding := 8 - (len(bitString) % 8)

		fmt.Println("Padded bits: ", padding)
		for i := 0; i < padding; i++ {
			bitString += "0"
		}
	}

	for i := 0; i < len(bitString); i += 8 {
		b, _ := strconv.ParseUint(bitString[i:i+8], 2, 8)
		res = append(res, byte(b))
	}
	fileIO.WriteToFile(res, outputFile)
}

func Decode(inputFile string, outputFile string, encoderMap map[rune]string){
	
	fmt.Println(encoderMap)

	// MAKE A REVERSE MAP STRING:RUNE
	decoderMap := make(map[string]rune)
	for k,v := range encoderMap{
		decoderMap[v] = k
		fmt.Println(v,string(k))
	}

	contents := fileIO.ReadFromFile(inputFile)

	var bitStream,decodedString strings.Builder

	end := "\nEND OF HEADER\n"
	t := 0
    // Read the file byte by byte
	for _,b := range contents{
		if t<15 {
			if end[t] == b{
				t++
			}else{
				t = 0
			}
		}
		if t>=15 {		
			bitStream.WriteString(fmt.Sprintf("%08b",b))
		}
	}

	// PERFORM THE DECODE OPERATION
	for i:=0;i<len(bitStream.String());i++{
		for j := i; j<len(bitStream.String());j++{
			if val, exists := decoderMap[bitStream.String()[i:j]]; exists{
				decodedString.WriteString(string(val))
				i = j
			}
		}
	}

	fileIO.WriteToFile(decodedString.String(),outputFile)
}