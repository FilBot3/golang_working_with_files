// Rename a file.
// Technically, this is just moving a file to rename it similar to how Linux's
// CoreUtils work using the mv command.
package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	log.Println("Moving/Renaming file.")
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
