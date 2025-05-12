// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
// Программа в первой строке получает на вход число n - количество чисел в последовательности,
//  во второй строке -- n чисел, входящих в данную последовательность.

package main

import "fmt"

func main() {
	var (
		n, res int
	)

	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	for i := 0; i < n; i++ {
		if array[i]%8 == 0 && isTwoDigit(array[i]) != 0 {
			res += array[i]
		}
	}

	fmt.Print(res)
}

func isTwoDigit(number int) int {
	digitCount := 0

	temp := number

	for temp > 0 {

		temp /= 10
		digitCount++

	}
	if digitCount == 2 {
		return number
	} else {
		return 0
	}
}
