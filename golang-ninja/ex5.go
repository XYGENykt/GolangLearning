package main

import "fmt"

func main() {
	number := 5
	var p *int
	p = &number
	fmt.Println(p)
	fmt.Println(&number)

	*p = 10
	fmt.Println(number)

	message := "hello"
	printMessage(&message)
	fmt.Println(message)

}

func printMessage(message *string) {
	*message += " world"
	fmt.Println(*message)
}
