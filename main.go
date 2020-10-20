package main

import (
	"fmt"
	"github.com/akhilerm/api-test/pkg/calc"
)

func main() {
	v := calc.NewVariables(12, 13)
	fmt.Printf("Sum = %d \n", v.Add())
}
