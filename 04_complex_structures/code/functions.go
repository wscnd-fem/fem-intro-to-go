package main

import "fmt"

func namedReturnPrintAge(age int) (ageOfW int, ageOfR int) {
	ageOfR = age + 21
	ageOfW = age + 20
	return
}

func spreadPrintAge(f func(a ...any) (int, error), ages ...int) {
	for _, value := range ages {
		f(value + 20)
	}
}

func fn[T any](f func(a T) T, b T) (T, error) {
	return f(b), nil
}

func funcao(a int) int {
	return a * 10
}

func main() {
	fmt.Println(namedReturnPrintAge(0))
	ageA, ageB := namedReturnPrintAge(0)

	fmt.Println("ageA:", ageA, "ageB:", ageB)
	spreadPrintAge(fmt.Println, 1, 2)

	x, y := fn(funcao, 2)

	fmt.Println("x:", x, "y:", y)
}
