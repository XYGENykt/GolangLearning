// Определите является ли данный символ цифрой или нет.

package main

import "fmt"

func main() {

	var (
		r rune
	)

	fmt.Scanf("%c", &r)

	if r >= '0' && r <= '9' {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
