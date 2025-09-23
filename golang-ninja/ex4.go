package main

import "fmt"

var msg string

func init() {
	msg = "from init"
}

func main() {
	fmt.Print(msg)
}
