package main

import "fmt"

func main() {

	aSentence := "sentence"

	for index, letter := range aSentence {
		if index%2 == 0 {
			fmt.Println("idx: ", index, "letter: ", string(letter))
		}
	}
}
