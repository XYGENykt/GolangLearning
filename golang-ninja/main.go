package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		message string
	)
	message = sayHello("Maksim", 15)
	print(message)
	print("2")
	result, err := enterTheClub(11)
	if err != nil {
		fmt.Println(err)
		return
		//log.Fatal(err)
	}
	fmt.Println(result)

}

func print(message string) {

	fmt.Println(message)
}

func sayHello(name string, age int) string {
	result := fmt.Sprintf("Привет, %s! Тебе %d лет!", name, age)
	return result
}

func enterTheClub(age int) (string, error) {
	if age < 18 {
		return fmt.Sprintf("Вам меньше 18 лет, вам %d!!!", age), errors.New("Вам нельзя")
	} else {
		return fmt.Sprintf("Вам можно входить так как вам %d!!!", age), nil
	}
}

func prediction(dayOfTheWeek string) string {

}
