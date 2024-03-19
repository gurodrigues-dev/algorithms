package main

import "fmt"

func quickSort(array []int) []int {

	if len(array) < 2 {
		return array
	}

	pivo := array[0]
	var menores []int
	var maiores []int

	for _, i := range array[1:] {
		if i <= pivo {
			menores = append(menores, i)
		}
	}

	for _, i := range array[1:] {
		if i > pivo {
			maiores = append(maiores, i)
		}
	}

	return append(append(quickSort(menores), pivo), quickSort(maiores)...)

}

func main() {

	x := quickSort([]int{10, 5, 3, 2, 8, 7, 4})

	fmt.Println(x)

}
