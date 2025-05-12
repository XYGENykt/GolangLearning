// Даны два натуральных числа. Выяснить, в каком из них больше цифр.
// Выведите цифру 1, если количество цифр в первом числе больше, чем количество цифр во втором числе.
// Выведите цифру 2, если количество цифр во втором числе больше, чем количество цифр в первом числе.
// Выведите цифру 0, если количество цифр в первом числе равно количеству цифр во втором числе.

package main

import (
	"fmt"
)

func countDigits(num int) int {
	if num == 0 {
		return 1
	}
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	digitsA := countDigits(a)
	digitsB := countDigits(b)

	if digitsA > digitsB {
		fmt.Println(1)
	} else if digitsB > digitsA {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}
