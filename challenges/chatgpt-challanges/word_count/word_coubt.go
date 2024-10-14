package main

import (
	"fmt"
	"strings"
)

func wordCount(word string) int {
	return len(strings.Fields(word))
}

func main() {
	fmt.Printf(`"Go is an awesome programming language" contém %v palavras`, wordCount("Go is an awesome programming language"))
}
