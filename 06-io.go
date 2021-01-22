package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open this file
	f, err := os.Open("06-io.go")
	if err != nil {
		// Note: Report errors on stderr
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		os.Exit(1)
	}
	defer f.Close() // Close file at end of function

	// Copy file contents to stdout; report outcome to stderr
	written, err := io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error copying to stdout:", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Success! %d bytes written\n", written)
}
