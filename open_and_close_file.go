// Opening and closing a file for editing text.
package main

import (
	"log"
	"os"
)

func main() {
	// Open a pointer to the file.
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Close our pointer to the file. This is needed for streams such as
	// network sockets and such.
	file.Close()

	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// Use these attributes individually or combined with an ORG for a second
	// argument of OpenFile()
	// e.g.
	//     os.O_CREATE|os.O_APPEND
	// or
	//     os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	//
	// os.O_RDONLY -> Read Only
	// os.O_WRONLY -> Write Only
	// os.O_RDWR   -> Read and Write
	// os.O_APPEND -> Append to the end of the file
	// os.O_CREATE -> Create if non-existant
	// os.O_TRUNC  -> Truncate file when opening.
}
