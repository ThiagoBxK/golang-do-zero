package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	var userA = User{Name: "Maria", Age: 42, Salary: 2792.50}
	var userB = User{"Lucas", 55, 5503.00}

	var pointer *User = &userA

	userA.Age = 43
	fmt.Println(userA, userB, pointer)
}
