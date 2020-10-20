package main

import (
	"fmt"
	"github.com/akhilerm/api-test/v2/pkg/calc"
)

func main() {
	v := calc.NewVariables(12, 13, 14)
	fmt.Printf("Sum = %d \n", v.Add())
}
