// Дана строка. Удалите k-ый символ в ней.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		x      string
		output string
	)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	fmt.Scan(&x)

	for i := 0; i < len(input); i++ {
		if i == x-1 {
			continue
		} else {
			output += string(input[i])
		}

	}

	fmt.Println(output)
}
