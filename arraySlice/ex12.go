// Дан массив, состоящий из целых чисел. Напишите программу, которая подсчитает количество элементов массива, которые меньше следующего за ним элемента.
package main

import "fmt"

func main() {
	var (
		n, x  int
		slice []int
		count int
	)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		slice = append(slice, x)
	}

	for i := 0; i < n; i++ {
		if i == n-1 {
			continue
		} else {
			if slice[i] < slice[i+1] {
				count++
			}

		}
	}

	fmt.Print(count)

}
