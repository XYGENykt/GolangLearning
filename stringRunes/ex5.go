// По данной строке определите, является ли она палиндромом? То есть, которое одинаково читается слева направо и справа налево. Например, слово "шалаш".

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		array  []string
		arrayP []string
		flag   bool
	)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for i := range input {
		array = append(array, string(input[i]))
	}

	for i := len(array) - 1; i >= 0; i-- {
		arrayP = append(arrayP, array[i])
	}

	for i := 0; i < len(array); i++ {
		if array[i] != arrayP[i] {
			flag = true
			break
		}
	}

	if flag == true {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}

}
