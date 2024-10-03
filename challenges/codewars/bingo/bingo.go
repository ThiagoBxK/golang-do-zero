package main

import "fmt"

// Desafio CodeWars : Bingo ( Or Not )

func bingo(numbers []int) string {
	winNumbers := []int{2, 9, 14, 7, 15}

	for _, winNumber := range winNumbers {
		found := false

		for _, number := range numbers {
			if number == winNumber {
				found = true
				break
			}
		}

		if !found {
			return "LOSE"
		}
	}

	return "WIN"
}

func main() {
	fmt.Println(bingo([]int{21, 13, 2, 7, 5, 14, 7, 15, 9, 10}))
}
