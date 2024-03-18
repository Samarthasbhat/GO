package main

import(
	"fmt"
)

func c(v string){
	switch v{
	case "John":
		fmt.Println("is an admin")
	case "Kat", "sally", "kevin":
		fmt.Println("Is an employee")
	
	}
}

func f(){
	switch c := 4; c{
	case 2:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("100")
	}
}

func main(){
	c("kevin")
	c("sally")
	f()
}