//
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a hard link.
	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create SymLink or Softlink
	log.Println("Creating SymLink")
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Get information on the link. If its a symlink, this will return
	// information about the Symlink rather than the file itself.
	// Hardlinks return information about the file they're linked to.
	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info %+v", fileInfo)

	// Change the ownership of the Symlink only.
	// Not the file it points to.
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
