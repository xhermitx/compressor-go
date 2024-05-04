package fileIO

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const HEADER_CONTENT_SEPARATOR = "\nEND OF HEADER\n"

func ReadFromFile(inputFile string) []byte{
	res,err := os.ReadFile(inputFile)
	if err!=nil{
		log.Fatal(err)
	}
	return res
}

// FUNCTION TO CREATE A FILE AND WRITE TO IT
func WriteToFile(fileContent interface{}, fileName string){

	f,err := os.OpenFile(fileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		log.Fatal(err)
	}
	defer f.Close()

	var data string

	switch v:= fileContent.(type){
	case string: data = v
	case []byte: data = string(v)
	default: 
		fmt.Println("content must be of type []byte or string")
	}
	_,err2 := f.WriteString(data)
	if err!=nil{
		log.Fatal(err2)
	}
}

func WriteHeader(fileName string,encoderMap map[rune]string){
	mapForJSON := make(map[rune]string)

	for k,v := range encoderMap{
		mapForJSON[k] = v
	}

	jsonData,err := json.Marshal(mapForJSON)
	if err != nil{
		log.Fatal(err)
	}

	err2 := os.WriteFile(fileName, jsonData, 0644)
	if err2 != nil{
		log.Fatal(err2)
	}

	WriteToFile(HEADER_CONTENT_SEPARATOR, fileName)
}

func ReadHeader(fileName string) map[rune]string{
	f, err := os.Open(fileName)	
	if err!=nil{
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var headerBuilder strings.Builder

	for scanner.Scan() {
        line := scanner.Text()

        // Check for the header separator line
        if line == "END OF HEADER" {
            break
        }

        // Write the line to the header string builder
        headerBuilder.WriteString(line)
    }

	headerJSON := headerBuilder.String()

    // Unmarshal the JSON header into a map (or a struct depending on your header structure)
    var headerMap map[rune]string // Using a map as an example
    err = json.Unmarshal([]byte(headerJSON), &headerMap)
    if err != nil {
        fmt.Println("Error unmarshalling JSON header:", err)
    }
	return headerMap
}