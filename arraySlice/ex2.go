// Напишите программу, которая переставляет соседние элементы массива (1-й элемент поменять с 2-м, 3-й с 4-м и т.д.
//
//	Если элементов нечетное число, то последний элемент остается на своем месте).
package main

import "fmt"

func main() {
	var (
		x, n  int
		slice []int
		temp  int
	)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		slice = append(slice, x)
	}

	for i := 1; i < n; {
		if len(slice)%2 == 0 {
			temp = slice[i]
			slice[i] = slice[i-1]
			slice[i-1] = temp
			i += 2
		} else {
			if i == n-1 {
				fmt.Print(slice[i])
				break
			}
			temp = slice[i]
			slice[i] = slice[i-1]
			slice[i-1] = temp
			i += 2
		}

	}

	for i := 0; i < n; i++ {
		fmt.Print(slice[i], " ")
	}
}
