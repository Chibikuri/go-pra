package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func main(){
	p := Person{
		Name: "Ryosuke",
		Age: 20,
	}
	fmt.Printf("first p: %+v\n", p)

	p2 := p
	p2.Name = "ryocchi"
	p2.Age = 21

	fmt.Printf("Is this rewritten? p: %+v\n", p)
	fmt.Printf("This is p2: %+v\n", p2)

	p3 := &p
	p3.Name = "yocchi"
	p3.Age = 23

	fmt.Printf("This should be rewritten p: %+v\n", p)
	fmt.Printf("This is p3: %+v\n", p3)
}
