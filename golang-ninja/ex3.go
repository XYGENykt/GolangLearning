package main

//замыкание
import "fmt"

func main() {
	// func() {
	// 	fmt.Println("Анонимная функция")
	// }()

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	inc2 := increment2()
	fmt.Println(inc2)
	fmt.Println(inc2)
	fmt.Println(inc2)
	fmt.Println(inc2)
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}
