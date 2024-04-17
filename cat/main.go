package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		// Handle input from stdin
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			printError(err)
			os.Exit(1)
		}
		return
	}
	// Process each file in the command line arguments
	for _, filename := range args {
		err := printFile(filename)
		if err != nil {
			// Print the error to stderr, matching the exact error output format
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
	}
}

// printFile opens and prints the content of the file to stdout
func printFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)
	return err
}

// printError writes an error message to stderr using the format specified
func printError(err error) {
	os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
}
