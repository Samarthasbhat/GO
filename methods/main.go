// 	
// func (receiver receiverType)funcName(arg argType) returnType {}

package main

import ("fmt")


type Bird struct {
	name string
}

func (b Bird)Fly() {
	fmt.Println(b.name, "is flying...")
}

func main(){
	b := Bird{"Raven"}
	b.Fly()
}