package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println("Error opening ")
	}
	defer file.Close()

	const bufferSize = 4
	b := make([]byte, bufferSize)

	for {
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		fmt.Println(string(b[:readTotal]))
	}
}
