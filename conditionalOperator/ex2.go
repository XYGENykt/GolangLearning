// Дано целое число
// n
// n. Выведите следующее за ним четное число.

// Входные данные
// Вводится натуральное число
// n
// n,  не большее
// 10000
// 10000.

// Выходные данные
// Выведите следующее за
// n
// n четное число.

// Sample Input 1:

// 15
// Sample Output 1:

// 16
// Sample Input 2:

// 20
// Sample Output 2:

// 22

package main

import "fmt"

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println(n + 2)
	} else {
		fmt.Println(n + 1)
	}
}
