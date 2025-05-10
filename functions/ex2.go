// Определите являются ли билеты счастливыми. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
// На вход подаются номера билетов - два шестизначных числа.
// Выведите 1, если они оба счастливые, в противном случае -1.

package main

import "fmt"

func main() {
	var (
		b1, b2 int
	)

	fmt.Scan(&b1, &b2)

	if checkLuckyB(b1) == 1 && checkLuckyB(b2) == 1 {
		fmt.Print(1)
	} else {
		fmt.Print(-1)
	}

}

func checkLuckyB(b int) int {
	var sum1, sum2, digit int

	for i := 0; i < 3; i++ {
		digit = b % 10
		b /= 10
		sum1 += digit
	}

	for i := 0; i < 3; i++ {
		digit = b % 10
		b /= 10
		sum2 += digit
	}

	if sum1 == sum2 {
		return 1
	} else {
		return 0
	}

}
