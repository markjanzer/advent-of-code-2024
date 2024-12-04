package lib

import (
	"io"
	"os"
)

const DataFileName string = "data.txt"

func GetDataString() string {
	return readDataFile(DataFileName)
}

func readDataFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read file content into a byte slice
	byteContent, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(byteContent)
}
