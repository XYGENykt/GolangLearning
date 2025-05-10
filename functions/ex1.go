// Дано два натуральных числа. Найдите количество разрядов каждого из них и выведите их произведение.
//Выведите произведение количества разрядов данных чисел.

package main

import "fmt"

func main() {
	var (
		x, y int
	)

	fmt.Scan(&x, &y)

	fmt.Print(digitCount(x) * digitCount(y))

}

func digitCount(x int) int {
	count := 0

	for x > 0 {
		x /= 10
		count++
	}
	return count
}
