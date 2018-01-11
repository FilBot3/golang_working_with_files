// Creating a blank file to use for demonstration purposes for Golang.
// Here we're Creating a file named test.txt, then checking if we had errors
// and then printing the file metadata, kind of. Not sure what it really is.
package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	log.Println("Createing new file")
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}
