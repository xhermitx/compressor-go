package fileIO

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const CHUNK_SIZE = 20000

func Encoder(fileContent []byte,encoderMap map[rune]string, charCount map[rune]int, fileName string) {

	// CALCULATE THE LENGTH OF THE ENCODED BIT STRING
	totalLength := 0
	for key, val := range encoderMap {
		totalLength += charCount[key] * len(val)
	}

	fmt.Println("Calculated length of the encoded bits : ", totalLength)

	encodedBits := ""
	extra := 0
	carryBits := ""

	for _,b := range fileContent{
		if extra > 0{
			encodedBits+=carryBits
			carryBits = ""
			extra = 0
		}
		encodedBits += encoderMap[rune(b)]

		if len(encodedBits) >= CHUNK_SIZE {
			extra = len(encodedBits)%CHUNK_SIZE
			carryBits = encodedBits[len(encodedBits)-extra:]
			encodedBits = encodedBits[:len(encodedBits)-extra]
			bitToByte(encodedBits, fileName)
			encodedBits = ""
		}
	}

	// fmt.Println("Total length of the encoded bits : ",len(encodedBits)+sum)
	bitToByte(encodedBits, fileName)
}

// CONVERTS THE ENCODED STRING OF BITS INTO A BYTE ARRAY TO WRITE TO A FILE
func bitToByte(bitString string, fileName string){

	var res []byte
	// CALCULATING THE PADDING BITS
	if len(bitString)%8 !=0{
		padding := 8-(len(bitString)%8)
		for i:=0;i<padding;i++{
			bitString += "0"
		}
	}

	for i:=0;i<len(bitString);i+=8{
		b,_ := strconv.ParseUint(bitString[i:i+8], 2, 8)
		res = append(res, byte(b))
	}

	writeToFile(res, fileName)
}

// FUNCTION TO CREATE A FILE AND WRITE TO IT
func writeToFile(fileContent []byte, fileName string){

	f,err := os.OpenFile(fileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		log.Fatal(err)
	}
	defer f.Close()

	_,err2 := f.WriteString(string(fileContent))
	if err!=nil{
		log.Fatal(err2)
	}
}