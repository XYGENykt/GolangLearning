// По данному числу определите, является ли оно палиндромом?
package main

import "fmt"

func main() {
	var (
		temp, input, digit, reverseInput int
	)
	fmt.Scan(&input)
	temp = input
	count := 0
	for temp > 0 {
		temp /= 10
		count++
	}
	temp = input
	reverseInput = 0
	for i := 0; i < count; i++ {
		if i > 0 {
			reverseInput *= 10
		}
		digit = temp % 10
		temp /= 10
		reverseInput += digit

	}
	if input == reverseInput {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
