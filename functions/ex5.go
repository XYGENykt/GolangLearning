// С помощью этой функции найти двойные факториалы трех целых чисел.

package main

import (
	"fmt"
)

// Функция для вычисления двойного факториала
func doubleFactorial(n int) int {
	result := 1
	for i := n; i > 0; i -= 2 {
		result *= i
	}
	return result
}

func main() {
	var numbers [3]int

	// Ввод трёх целых чисел
	for i := 0; i < 3; i++ {
		fmt.Scan(&numbers[i])
	}

	// Вычисление и вывод двойного факториала для каждого числа
	for _, num := range numbers {
		fmt.Print(doubleFactorial(num), " ")
	}
}
