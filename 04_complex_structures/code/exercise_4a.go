package main

import "fmt"

func average(a, b, c float32) float32 {
	return (a + b + c) / 3
}

func main() {
	average := average(1.2, 2.2, 3.5)
	fmt.Println(average)
}
