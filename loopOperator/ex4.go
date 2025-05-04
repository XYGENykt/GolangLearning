package main

import "fmt"

func main() {
	var x, a int
	fmt.Scan(&x)
	count := 0
	for i := 1; i <= x; i++ {
		fmt.Scan(&a)
		a = a % 10
		if a == 0 {
			count++
		}

	}
	fmt.Println(count)
}
