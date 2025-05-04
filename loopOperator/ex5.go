// Сколько раз делится. По данному натуральному числу a, определите, сколько раз a делится на 3 нацело (без остатка).

package main

import "fmt"

func main() {
	var a, count int
	fmt.Scan(&a)

	count = 0
	for a > 0 {
		if a%3 == 0 {
			count++
			a /= 3

		} else {
			break
		}

	}
	fmt.Println(count)

}
