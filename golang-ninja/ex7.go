package main

import "fmt"

func main() {

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	getGameReuslt(&results)
	fmt.Println(results)
	calculateScore(results)
}

func getGameReuslt(res *[]string) []string {
	var s string
	fmt.Scan(&s)

	if s == "w" {
		*res = append(*res, s)
	} else if s == "d" {
		*res = append(*res, s)
	} else if s == "l" {
		*res = append(*res, s)
	} else {
		fmt.Println("invalid input")
		return *res
	}
	return *res
}

func calculateScore(input []string) {
	score := 0
	for i := range input {
		if input[i] == "w" {
			score += 3
		} else if input[i] == "d" {
			score += 1
		} else {
			score += 0
		}
	}
	fmt.Println(score)
}
