package main

import (
	"fmt"
)

type User struct {
	Name    string
	Surname string
	Age     int
	Salary  float64
}

// Receptor com ponteiro u de User
func (u *User) fullName() string {
	return fmt.Sprintln(u.Name, u.Surname)
}

func main() {
	userA := User{
		Name:    "Carlos",
		Surname: "Silva",
		Age:     28,
		Salary:  4750.00,
	}

	fmt.Println(userA.fullName())
}
