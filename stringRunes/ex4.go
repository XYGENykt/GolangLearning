// Дана строка. Определите, какой символ в ней встречается раньше: 'x' или 'w'. Если таких символов нет, вывести "-1".

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

	countX := int(strings.Count(input, "x"))
	countXIndex := int(strings.Index(input, "x"))
	countW := int(strings.Count(input, "w"))
	countWIndex := int(strings.Index(input, "w"))

	if countX > countW {
		fmt.Print("x")
	} else if countX < countW {
		fmt.Print("w")
	}

	if countX == countW && countXIndex < countWIndex {
		fmt.Print("x")
	} else if countX == countW && countXIndex > countWIndex {
		fmt.Print("w")
	}

	if countX == 0 && countW == 0 {
		fmt.Print("-1")
	}

}
