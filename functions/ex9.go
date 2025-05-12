// Даны два предложения. Найти общее количество букв 'b' в них.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLetterB(s string) int {
	count := 0
	for _, char := range strings.ToLower(s) {
		if char == 'b' {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение первой строки
	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)

	// Чтение второй строки
	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)

	// Подсчёт букв 'b' в каждой строке
	count1 := countLetterB(line1)
	count2 := countLetterB(line2)

	// Общее количество
	total := count1 + count2

	fmt.Print(total)
}
