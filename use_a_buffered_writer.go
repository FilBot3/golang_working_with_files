//
package main

import (
	"bufio"
	"log"
	"os"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Open a file for writing.
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	checkerr(err)
	defer file.Close()

	// Create a buffered writer from the file.
	bufferedWriter := bufio.NewWriter(file)

	// Write bytes to the buffer
	bytesWriteen, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	checkerr(err)
	log.Printf("Bytes written: %d\n", bytesWritten)

	// Write string to buffer
	// Also available are WriteRune() and WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered String\n",
	)
	checkerr(err)
	log.Printf("Bytes written: %d\n", bytesWritten)

	// Check how much is stored in buffer waiting
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// See how much buffer is available
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// Write memory buffer to disk.
	bufferedWriter.Flush()

	// Revert any changes done to the buffer that have not yet been written to
	// file with Flush().
	// We just flushed, so there are no changes to revert. The writer that you
	// pass as an argument is where the buffer will output to.
	// If youw ant to change to a new writer.
	bufferedWriter.Reset(bufferedWriter)

	// See how much buffer is available
	bytesAvailable = bufferedWrite.Available()
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// Resize the buffer. The first argument is a writer where the buffer should
	// output to. In this case we are using the same buffer. If we chose a
	// number that was smaller than the existing buffer, like 10, we would not
	// get back a buffer of size 10, we will get back a buffer the size of the
	// original since it was already enough (default 4096)
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// Check available buffer size after resizing.
	bytesAvailable = bufferedWriter.Available()
	log.Printf("Available buffer: %d\n", bytesAvailable)
}
