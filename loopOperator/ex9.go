// Из данного числа выбросите цифры 5 и 7 , при этом порядок остальных цифр не меняется.
package main

import "fmt"

func main() {
	var (
		input int
	)
	fmt.Scan(&input)
	fmt.Println(removeDigits(input))

}
func removeDigits(n int) int {
	result := 0
	multiplier := 1

	for n > 0 {
		digit := n % 10
		n = n / 10
		if digit != 5 && digit != 7 {
			result += digit * multiplier
			multiplier *= 10
		}
	}

	return result
}
