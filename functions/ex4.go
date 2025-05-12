// Дано два натуральных числа, не заканчивающиеся на 0. Замените каждое число на обратное и вычислите их сумму.
// Например, дается два числа 624 и 1024. Каждое заменяем на обратное, то есть 624⇒426,1024⇒4201. Затем находим их сумму: 426+4201=4627.

package main

import "fmt"

func main() {
	var (
		x, y int
	)

	fmt.Scan(&x, &y)

	fmt.Print(reverse(x) + reverse(y))

}

func reverse(number int) int {
	var (
		digitCount, digit, res int
	)
	temp := number

	for temp > 0 {
		temp /= 10
		digitCount++
	}

	temp = number
	for i := 0; i < digitCount; i++ {
		digit = temp % 10
		temp /= 10

		if i == digitCount-1 {
			res = (res * 10) + digit
		} else {
			res = (res * 10) + digit

		}

	}

	return res

}
