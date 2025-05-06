// Дан массив, состоящий из целых чисел. Напишите программу, которая уменьшает все элементы массива на минимальный элемент.

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
	min := array[minIndex]

	for i := range array {
		if array[i] < min {
			min = array[i]
			minIndex = i
		}
	}

	for i := range array {
		array[i] -= min
		fmt.Print(array[i], " ")
	}
}
