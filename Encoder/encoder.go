package encoder

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var line string

    // Read line by line until the header is found
    for {
        line, err = reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading file:", err)
            return
        }

        // Check if the line is the end of the header
        if strings.Contains(line, "END OF HEADER") {
            break
        }
    }

	var bitStream,decodedString strings.Builder
	carry := ""

    // Now, read the rest of the file byte by byte
    for {
        b, err := reader.ReadByte()
        if err != nil {
            if err == io.EOF {
                break // End of file
            }
            fmt.Println("Error reading byte:", err)
            return
        }

		if len(carry)>0{
			bitStream.WriteString(carry)
			carry = ""
		}

		bitStream.WriteString(fmt.Sprintf("%08b",b))

		if len(bitStream.String()) >= CHUNK_SIZE{
			decodedString,carry = decodeChunk(bitStream, decoderMap)
			fileIO.WriteToFile(decodedString.String(), outputFile)
			bitStream.Reset()
		}
    }

	bitStream.WriteString(carry)
	decodedString,carry = decodeChunk(bitStream,decoderMap)
	fileIO.WriteToFile(decodedString.String(), outputFile)
}

func decodeChunk (bitStream strings.Builder, decoderMap map[string]rune) (strings.Builder,string){
	var decodedString strings.Builder
	// PERFORM THE DECODE OPERATION
	t := 0
	for i:=0;i<len(bitStream.String());i++{
		for j := i; j<len(bitStream.String());j++{
			if val, exists := decoderMap[bitStream.String()[i:j]]; exists{
				decodedString.WriteString(string(val))
				i = j
				t = j
			}
		}
	}
	return decodedString, bitStream.String()[t:]
}