package main

import (
	"fmt"
)

func main() {
	fileName := "/main.txt"

	file := createFile(fileName)
	defer file.Close()

	if file == nil {
		fmt.Println("Falha ao criar o arquivo.")
		return
	}

	writeFile(file, "Testando...")
}
