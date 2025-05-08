// Дан список положительных целых чисел. Найдите все числа кратные 3 и при этом оканчивающиеся на 7,
// и замените каждое из этих чисел на общее коллчество таких чисел в массиве.

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

	for i := range array {
		if array[i]%3 == 0 && array[i]%10 == 7 {
			count++
		}
	}

	for i := range array {
		if array[i]%3 == 0 && array[i]%10 == 7 {
			array[i] = count
		}
	}

	for i := range array {
		fmt.Print(array[i], " ")
	}

}
