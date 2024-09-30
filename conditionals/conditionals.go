package main

import "fmt"

func main() {
	var name string = "admin"
	var token string = "API_TOKEN"
	var expectedToken string = "admin"

	// If & Else
	if token == expectedToken {
		fmt.Println("Autenticado!")
	} else {
		fmt.Println("Login Invalido!")
	}

	// Switch
	switch {
	case name == "Admin":
		fmt.Println("Olá, Chefe!")
	default:
		fmt.Println("Olá", name)
	}
}
