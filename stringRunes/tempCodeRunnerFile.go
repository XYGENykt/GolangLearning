// По данной строке, найдите ее k-ый символ. Выведите k-ый символ строки, если он существует, в противном случае выведите "NO".
// BeeGeek
// 4
// G

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		c []rune
		k int
	)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	c = []rune(input)

	fmt.Scan(&k)

	if k <= len(c) {
		fmt.Println(string(c[k-1]))
	} else {
		fmt.Println("NO")
	}

}
