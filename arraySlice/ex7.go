// Дан массив, состоящий из целых чисел. Напишите программу, которая выведет разность индексов максимального и минимального элементов массива. Индексация начинается с нуля.

package main

import "fmt"

func min(array []int) int {

	minIndex := 0
	min := array[minIndex]
	for i := range array {
		if array[i] < min {
			min = array[i]
			minIndex = i
		}
	}
	return minIndex
}

func max(array []int) int {

	maxIndex := 0
	max := array[maxIndex]
	for i := range array {
		if array[i] > max {
			max = array[i]
			maxIndex = i
		}
	}
	return maxIndex
}

func main() {
	var (
		count int
	)
	fmt.Scan(&count)
	array := make([]int, count)

	for i := range array {
		fmt.Scan(&array[i])
	}

	fmt.Println(max(array) - min(array))

}
