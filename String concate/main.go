package main

import(
	"fmt"
	"strings"
	"bytes"
)

func main() {
	name:= "John"+" "+"Doe"
	fmt.Printf("Plus + operator %v", name)
	fmt.Println()
	fmt.Println("It","works!")

	// String concatenation using string append

	u:="This"
	v := " is string concatenation"
	u += v
	fmt.Printf("%v",u)

	fmt.Println()
	// String join() function to concatenate two strings

	s := []string{"This","is", "a", "string."}

	vs:= strings.Join(s, " ")

	fmt.Printf("%v",vs)

	fmt.Println()
	//The Sprintf() method to concatenate strings in Go

	s1 := "abc"
	s2 := "def"

	sprint := fmt.Sprintf("%s%s",s1,s2)
	
	fmt.Println(sprint)

	var b bytes.Buffer

	b.WriteString("def")
	b.WriteString("ghi")

	fmt.Println(b.String())

	//  Repeat() Method to concatenate same string multiple times
	fmt.Println(strings.Repeat("abc",3))
	
}