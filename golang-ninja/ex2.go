package main

import "fmt"

func main() {
	fmt.Println((findMin(1, 323, 4123, -123, 34, 3)))
	fmt.Println((findMax(1, 323, 4123, -123, 34, 3)))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}
