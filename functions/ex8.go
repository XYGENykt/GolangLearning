// Даны два натуральных числа. Выяснить, в каком из них сумма цифр больше.
// Выведите цифру 1, если сумма цифр первого числа больше, чем сумма цифр второго числа.
// Выведите цифру 2, если сумма цифр второго числа больше, чем сумма цифр первого числа.
// Выведите цифру 0, если сумма цифр первого числа равно сумме цифр второго числа.

package main

import (
	"fmt"
)

func sumDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	sumA := sumDigits(a)
	sumB := sumDigits(b)

	if sumA > sumB {
		fmt.Println(1)
	} else if sumB > sumA {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}
