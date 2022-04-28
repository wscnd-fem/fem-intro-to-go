package main

import "fmt"

func variadicAverage(a ...float32) (float32, float32) {
	var sum, sum2, size float32 = 0, 0, float32(len(a))

	for _, v := range a {
		sum += v
	}

	for i := 0; i < len(a); i++ {
		sum2 += a[i]
	}

	return sum / size, sum2 / size
}

func main() {
	fmt.Println(variadicAverage(1.2, 2.2, 3.5))
}
