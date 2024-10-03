package main

import "fmt"

func main() {
	// Variavel padrão
	var name = "Sofia"

	// Variável curta com tipo inferido
	surname := "Anderson"

	// Declaração de múltiplas variáveis
	var (
		age    int     = 35
		height float32 = 1.75
	)

	work, country, salary := "Programmer", "Georgia", 12500.00
	message := fmt.Sprintf("They are %d years old, %.2f meters tall, work as an %s, live in the %s, and earn $%.2f per year.", age, height, work, country, salary)

	fmt.Println(`Hello`, name, surname, "welcome to Go!")
	fmt.Println(message)

}
