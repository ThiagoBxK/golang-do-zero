package main

import (
	"fmt"
	"os"
	"path"
)

func getDirname() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", fmt.Errorf("Erro ao obter o diretório: %w", err)
	}

	return dir, nil
}

func createFile(fileName string) *os.File {
	dirname, err := getDirname()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	filesDir := path.Join(dirname, "/os/files/")
	filePath := path.Join(filesDir, fileName)

	file, err := os.Create(filePath)

	if err != nil {
		fmt.Println("Erro ao criar o arquivo %w", err)
		return nil
	}

	fmt.Println("Arquivo criado com sucesso!")
	return file
}

func writeFile(file *os.File, content string) {
	_, err := file.WriteString(content)

	if err != nil {
		fmt.Println("Erro ao escrever o arquivo %w", err)
		return
	}

	fmt.Println("Arquivo escrito com sucesso!")
}
