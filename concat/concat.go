package main

import (
	"fmt"
)

func main() {

	var (
		name    string = "Isabelle"
		message string = fmt.Sprintf("Hello, %s", name)
	)

	fmt.Println("Hello,", name)
	fmt.Println("Hello, " + name)
	fmt.Println(message)
}
