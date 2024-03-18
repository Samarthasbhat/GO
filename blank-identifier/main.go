package main

import( 
	"fmt"
)

func f() (int, int){
	
	return 42, 43
}

func f1() int {
	return 900
}

func main(){
	v, _:= f()
	fmt.Printf("The type of v :%T\n", v)
	fmt.Println(v)

	r:= f1()
	_ = r + 10

	fmt.Println(r)
}
