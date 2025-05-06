// Дан массив, состоящий из целых чисел. Напишите программу, которая выведет индекс минимального элемента массива. Индексация начинается с нуля.

package main

import "fmt"

func main() {
	var (
		count int
	)
	fmt.Scan(&count)
	array := make([]int, count)

	for i := range array {
		fmt.Scan(&array[i])
	}
	minIndex := 0

	for i := range array {
		if array[i] < array[minIndex] {
			minIndex = i

		}
	}
	fmt.Println(minIndex)

}
