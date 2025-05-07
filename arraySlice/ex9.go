// Дан массив, состоящий из целых чисел. Напишите программу, которая определяет, есть ли в массиве ХОТЯ БЫ два одинаковых элемента.
// Необходимо вывести "YES", если  в массиве существуют ХОТЯ БЫ два одинаковых элемента. В противном случае следует вывести "NO".

package main

import "fmt"

func main() {
	var (
		count, n int
	)
	fmt.Scan(&n)
	array := make([]int, n)

	for i := range array {
		fmt.Scan(&array[i])
	}

	for i := range array {
		for k := range array {
			if array[k] == array[i] {
				count++
			}
		}
		if count > 1 {
			fmt.Print("YES")
			break
		} else {
			count = 0
		}
	}
	if count == 0 {
		fmt.Print("NO")

	}
}
