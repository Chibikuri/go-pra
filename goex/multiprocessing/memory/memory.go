package main

import "fmt"

func main() {
	var data int
	go func() { data++ }()
	if data == 0 {
		fmt.Println("The value is 0.")
	} else {
		fmt.Printf("The value is %v.\n", data)
	}
}
