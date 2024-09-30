package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

type Users map[string]User

func main() {
	users := Users{
		"Matheus": User{
			Name:   "Matheus",
			Age:    32,
			Salary: 3450.00,
		},
	}

	users["Carlos"] = User{
		Name:   "Carlos",
		Age:    27,
		Salary: 5450.00,
	}

	value, exists := users["João"]

	if !exists {
		fmt.Println("Jõao não fez a prova!", value)
	}

	fmt.Println(users)

}
