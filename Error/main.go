package main

import (
  "fmt"
  "os"
)

func Open(name string) (*os.File, error) {
  f, err := os.Open(name)
  if err != nil {
    return nil, err // Return nil for file and the error
  }
  return f, nil // Return the opened file and nil for no error
}

func main() {
  file, err := Open("name.txt")
  if err != nil {
    fmt.Println("Error opening file:", err)
    // Handle the error here, maybe exit or retry
  } else {
    // Use the opened file 'file' here
    defer file.Close() // Close the file when the function exits
    fmt.Println("File opened successfully!")
  }
}
