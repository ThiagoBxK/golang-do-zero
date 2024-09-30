package main

import "fmt"

func counter() func(increment int) int {
	var counter int = 0

	return func(increment int) int {
		counter += increment
		return counter
	}
}

func main() {
	var count = counter()

	count(1)
	count(1)
	count(1)

	fmt.Println(count(1))
}
