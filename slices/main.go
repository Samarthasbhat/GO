package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := make([]byte, 3) //creates a new byte slice with an initial length of 3
	data := []byte{4, 5, 6}
	result := Append(slice, data)
	fmt.Println(result)

	slice1 := []int{1,2,3}
	slice2 := []int{1,2,3}

	// fmt.Println( slice1 == slice2)  // use reflect package 
	result1 := reflect.DeepEqual(slice1, slice2)
	fmt.Println(result1) 
}

func Append(slice, data []byte) []byte{
	l:= len(slice)
	if l + len(data) > cap(slice){
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:l+len(data)]
	copy(slice[l:], data)
	return slice
}