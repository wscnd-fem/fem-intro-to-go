package main

import "fmt"

var outsideVariable = "outside"

func main() {
	var varWithType string = "Beyonce"
	fmt.Println(varWithType)

	var inferVarType = "Inferred Type"
	fmt.Println(inferVarType)

  fmt.Println(outsideVariable)

	var noActualValue string
	fmt.Println(noActualValue)

	var allocatedValue int
	fmt.Println(allocatedValue)

	var multiple, variables = "valor1", 1
	fmt.Println(multiple, variables)

	mostCommon := "Most common"
  fmt.Println(mostCommon)
}
