// Дана строка. Найдите перевернутую ей строку.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	for i := len(input) - 1; i >= 0; i-- {
		fmt.Print(string(input[i]))
	}

}
