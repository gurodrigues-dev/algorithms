package main

import (
	"fmt"
	"sort"
)

func twoSum(numbers []int, target int) []int {
	sort.Ints(numbers)
	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return []int{}
}

func main() {
	resultado := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(resultado)
}
