package main

import "fmt"

func main() {
	token := "ADMIN_TOKEN"

	// Usando if & else
	if token == "ADMIN_TOKEN" {
		fmt.Println("Autenticado!")
	} else {
		fmt.Println("Login Invalido!")
	}

	// Usando Switch
	switch {
	case token == "ADMIN_TOKEN":
		fmt.Println("Autenticado!")
	default:
		fmt.Println("Login Invalido!")
	}
}
