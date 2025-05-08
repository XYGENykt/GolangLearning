// Дан список, упорядоченный по не убыванию элементов в нем. Определите, сколько в нем различных элементов.
// Если все элементы массива одинаковы, то количество различных элементов равняется 1.

package main

import "fmt"

func main() {
	var (
		n int
	)

	fmt.Scan(&n)

	array := make([]int, n)

	for i := range array {
		fmt.Scan(&array[i])
	}

	count := 0

	if n > 1 {
		count++
		for i := 0; i < n-1; i++ {
			if array[i] != array[i+1] {
				count++
			}

		}

	}
	if n == 1 {
		count++
	}

	fmt.Print(count)
}
