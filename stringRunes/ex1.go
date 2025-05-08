// Дана строка, содержащая пробелы. Найдите, сколько в ней слов.

// Слово – это последовательность не пробельных символов. Слова разделены одним пробелом, первый и последний символ строки – не пробел.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаём читатель для ввода с консоли
	reader := bufio.NewReader(os.Stdin)

	// Читаем строку до символа переноса строки
	input, _ := reader.ReadString('\n')

	// Удаляем символ переноса строки в конце
	input = strings.TrimSpace(input)

	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Выводим результат
	fmt.Println(len(words))
}
