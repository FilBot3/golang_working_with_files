// Delete a file from the system.
package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Removing file.")
	err := os.Remove("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
}
