package main

import "fmt"

func binarySearch(array []int, target int) int {

	smaller := 0
	bigger := len(array) - 1

	for smaller <= bigger {
		mid := (smaller + bigger) / 2
		guess := array[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			bigger = mid - 1
		}

		if guess < target {
			smaller = mid + 1
		}
	}

	return -1

}

func main() {

	array := []int{2, 4, 6, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 69, 82, 93, 119}

	index := binarySearch(array, 82)

	fmt.Println(index)

}
