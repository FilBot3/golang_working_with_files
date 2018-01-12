//
package main

import (
	"io"
	"log"
	"os"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Open the original file
	originalFile, err := os.Open("test.txt")
	checkerr(err)
	defer originalFile.Close()

	// Create new file.
	newFile, err := os.Create("test_copy.txt")
	checkerr(err)
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	checkerr(err)
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents and flush memory to disk to "save"
	err = newFile.Sync()
	checkerr(err)
}
