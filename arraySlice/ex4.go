// Напишите программу, которая циклически сдвигает элементы массива вправо. Например, если элементы нумеруются,
// начиная с нуля, то 0-й элемент становится 1-м, 1-й становится 2-м, ..., последний становится 0-м, то есть массив
// [3,5,7,9] превращается в массив [9,3,5,7]

package main

import "fmt"

func main() {
	var (
		n, x, temp int
		array      []int
	)

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		array = append(array, x)
	}

	temp = array[n-1]
	for i := n - 1; i > 0; i-- {
		array[i] = array[i-1]

	}

	array[0] = temp

	for i := 0; i < n; i++ {
		fmt.Print(array[i], " ")
	}
}
