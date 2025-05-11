// Даны две последовательности. Первая: 1,2,3,...,n, а вторая: 1,2,3,...,m. Найдите средние арифметические обеих последовательностей и выведите их сумму.

package main

import "fmt"

func main() {
	var (
		x, y int
	)

	fmt.Scan(&x, &y)

	fmt.Print(average(x) + average(y))

}

func average(number int) float64 {
	var res float64
	for i := 1; i <= number; i++ {
		res += float64(i)
	}
	res /= float64(number)

	return res

}
