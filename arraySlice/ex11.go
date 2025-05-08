// Дано целое число N и набор из N целых чисел. Вывести все четные числа из данного набора и кол-во таких чисел.
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
		if array[i]%2 == 0 && array[i] != 0 {
			fmt.Print(array[i], " ")
			count++
		}
	}
	fmt.Println()
	fmt.Print(count)

}
