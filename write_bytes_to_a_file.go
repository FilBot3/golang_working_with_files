//
package main

import (
	"log"
	"os"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Open a new file for writing only
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	checkerr(err)
	defer file.Close()

	// Write bytes to file
	byteSlice := []byte("Bytes!\n")
	byteWritten, err := file.Write(byteSlice)
	checkerr(err)
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
