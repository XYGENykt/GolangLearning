// Измените регистр символа, если он был латинской буквой: сделайте его заглавным, если он был строчной буквой и наоборот.

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var (
		c rune
	)

	fmt.Scanf("%c", &c)

	if c >= 65 && c <= 90 {
		fmt.Print(string(unicode.ToLower(c)))
	}
	if c >= 97 && c <= 122 {
		fmt.Print(string(unicode.ToUpper(c)))
	}
}
