package main

import "fmt"

func main(){
	name := "Ryosuke"
	fmt.Printf("name: %v\n", name)

	namePoint := &name

	fmt.Printf("namePoint: %v\n", namePoint)

	fmt.Printf("namePoint: %v\n", *namePoint)

	*namePoint = "yocchi"

	fmt.Printf("*namePoint after input yocchi:%v\n", *namePoint)

	fmt.Printf("namePoint after input yocchi:%v\n", namePoint)

	fmt.Printf("name after input yocchi:%v\n", name)
}
