package main

import "fmt"

func main(){
	name := "Ryosuke"
	fmt.Printf("name :%v\n", name)

	//ここにポインタを渡している
	namePoint := &name

	//ポインタの値を参照したい場合は*をつける必要がある
	fmt.Printf("namePoint :%v\n", namePoint)
	fmt.Printf("namePoint :%v\n", *namePoint)
}
