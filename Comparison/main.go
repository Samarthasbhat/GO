package main

import (
	"fmt"
	"strings"
)


func main() {
	a := "Hello"
	b := "Hi"

// Double equals operator
	check := a==b

	fmt.Println(check)

	check1 := a != b

	fmt.Println(check1)

	// General comparison operators

	s1 := "abc"
    s2 := "cab"
         
    v := s1 > s2
    v1 := s1 < s2
    v2 := s1 >= s2
    v3 := s1 <= s2
 
    fmt.Println(v)  // false
    fmt.Println(v1)  // true
    fmt.Println(v2)  // false
    fmt.Println(v3)   // true

	v4 := strings.Compare(s1, s2) // -1

	fmt.Println(v4)  // As you can see, it returns o if they are the same. If the first one is greater then it returns 1. If it’s lesser then it returns -1. The greater and lesser values are defined by their lexicographical order

	// Strings EqualFold() method to compare strings

	v5 := strings.EqualFold(s1,s2) // cases don't matter

	fmt.Println(v5) // true
}