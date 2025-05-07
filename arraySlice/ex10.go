// Дан список. Выведите те его элементы, которые встречаются в списке только один раз.

package main

import "fmt"

func main() {
	var (
		n, count    int
		arrayResult []int
	)
	fmt.Scan(&n)
	array := make([]int, n)

	for i := range array {
		fmt.Scan(&array[i])
	}

	for k := range array {
		for i := range array {
			if array[k] == array[i] {
				count++
			}

		}
		if count == 1 {
			arrayResult = append(arrayResult, array[k])
		}
		count = 0

	}

	for i := range arrayResult {
		fmt.Print(arrayResult[i], " ")
	}
}
