package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		name := os.Args[1]
		file, err := os.Open(name)
		if err != nil {
			fmt.Printf("Error occured.Error is : %v\n", err.Error())
		} else {
			fileInfo, _ := file.Stat()
			fileSize := fileInfo.Size()
			text := make([]byte, fileSize)
			file.Read(text)
			fmt.Print(string(text))
			file.Close()
		}
	}
}
