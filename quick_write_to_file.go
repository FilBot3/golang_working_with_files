// Quickly dump a Slice of Bytes into a file.
// This is handy for dumping small amounts of files or full files.
// Not so useful for writing to a log file.
package main

import (
	"io/ioutil"
	"log"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	checkerr(err)
}
