// Дан массив, состоящий из целых чисел. Напишите программу, которая определяет, есть ли в массиве пара соседних элементов с одинаковыми знаками.
package main

import "fmt"

func main() {
	var (
		x, n, minus, plus int
		slice             []int
	)
	srav := false
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		slice = append(slice, x)
	}

	for i := 0; i < n; i++ {
		if otric(slice[i]) == 1 {
			minus++
			plus = 0
		} else {
			plus++
			minus = 0
		}
		if minus == 2 || plus == 2 {
			fmt.Print("YES")
			srav = true
			break
		}
	}
	if srav == false {
		fmt.Print("NO")
	}

}

func otric(g int) int {
	if g < 0 {
		return 1
	} else {
		return 0
	}
}
