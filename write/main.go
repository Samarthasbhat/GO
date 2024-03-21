package main

import (
	// "fmt"
	// "os"
	"io/ioutil"
)

// func main(){

// 	f, err := os.Create("text.txt")
// 	if err != nil{
// 		fmt.Println(err)
// 	}
// 	defer f.Close()

	
// }

func main() {
    s := []byte("This is a string")                // convert string to byte slice
    ioutil.WriteFile("testfile.txt", s, 0644)      // the 0644 is octal representation of the filemode
}