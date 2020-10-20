package main

import (
	"fmt"
	"github.com/akhilerm/api-test/pkg/calc"
	calcv2 "github.com/akhilerm/api-test/v2/pkg/calc"
)

func main() {

	v1 := calc.NewVariables(12, 13)
	fmt.Printf("Sum = %d \n", v1.Add())

	v2 := calcv2.NewVariables(12, 13, 14)
	fmt.Printf("Multiply = %d \n", v2.Multiply())
}
