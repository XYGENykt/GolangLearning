// Вводится ненормированная строка, у которой может быть более одного пробела между словами. Привести ее к нормированному виду, т.е.
// между словами оставить только один пробел. Гарантируется, что в начале и в конце строки нет пробелов.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		array []string
	)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	array = append(array, string(input[0]))
	for i := 1; i < len(input); i++ {

		if string(input[i-1]) == " " && string(input[i]) == " " {
			continue
		} else {
			array = append(array, string(input[i]))
		}

	}

	for i := range array {
		fmt.Print(array[i])
	}
}
