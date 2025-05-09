// Дана строка. Известно, что она содержит ровно две одинаковые буквы. Найдите эти буквы. Гарантируется, что повторяются буквы только одного вида.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		array           []string
		count, countMax int
		charPopular     string
	)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	for i := range input {
		array = append(array, string(input[i]))
	}

	for i := range array {
		for k := 0; k < len(array); k++ {
			if array[i] == array[k] {
				count++
			}
		}
		if count > countMax {
			countMax = count
			charPopular = array[i]
		}
		count = 0
	}
	fmt.Print(charPopular)

}
