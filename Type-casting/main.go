package  main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var a int = 42
	f := float64(a)
	fmt.Println(f)

	fmt.Println("Type of f:", reflect.TypeOf(f))

	// strconv

	var s string = "42"
	v, _ := strconv.Atoi(s)

	fmt.Println(v)

	var i int = 42
	str := strconv.Itoa(i)
	
	fmt.Println(str)

	fl := 12.34567
	il := int(fl)
	fmt.Println(il)

	ii := 34
	ff := float64(ii)

	fmt.Println(ff)

	// Strings and bytes conversion

	var s1 string = "Hello World"
	var b []byte = []byte(s1)


	fmt.Println(b)

	ss := string(b)
	fmt.Println(ss)

	// Type conversion during conversion
	aa := 6/3
	f1 := 6.3/3

	fmt.Println(aa, f1)

}