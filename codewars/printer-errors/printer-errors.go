package main

import (
	"fmt"
	"regexp"
)

// Desafio CodeWars : Printer Errors

func printerError(str string) string {
	regex := regexp.MustCompile("[^a-mA-M]")
	matches := regex.FindAllString(str, -1)

	return fmt.Sprintf("%d/%d", len(matches), len(str))
}

func main() {
	str := "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"
	fmt.Println(printerError(str))
}
