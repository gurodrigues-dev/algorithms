package main

import "fmt"

func calcFat(x int) int {
	if x == 1 {
		return x
	}

	return x * calcFat(x-1)
}

func main() {
	x := calcFat(3)
	fmt.Println(x)
}
