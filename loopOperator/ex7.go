// Количество цифр 7
// Посчитать, сколько раз встречается цифра 7 в последовательности чисел от 1 до N, включая N
package main

import (
	"fmt"
)

func main() {
	var n, k int
	count := 0
	digit := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {

		for k = i; k > 0; {
			digit = k % 10
			if k == 7 || digit == 7 {
				count++
			}
			k = k / 10
		}

	}
	fmt.Println(count)

}
