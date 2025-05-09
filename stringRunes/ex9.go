// По заданной строчной букве латинского алфавита, выведите все строчные буквы латинского алфавита, начиная от начала до заданной буквы.

package main

import "fmt"

func main() {

	var (
		c rune
	)

	fmt.Scanf("%c", &c)

	for i := 'a'; i <= c; i++ {
		if c >= i {
			fmt.Print(string(i), " ")
		}
	}
}
