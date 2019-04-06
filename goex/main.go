package main

import (
	"fmt"

	"./qubits"
	circuit "github.com/Chibikuri/lambdaQ/internal/circuit"
)

func main() {
	fmt.Println(qubits.Exconst(1))
	circuit.Exconst(3)
}
