package main

import "fmt"

// Se o parâmetro fosse passado como uma variável normal, o valor original não seria alterado,
// Pois a função receberia uma cópia da variável, e qualquer modificação ocorreria apenas na cópia.
func changeName(pointer *string) {
	// Mudando valor via ponteiro
	*pointer = "James"
}

func main() {
	name := "Alice"

	// Endereço de memoria da variavel
	pointer := &name

	fmt.Printf("Ponteiro: %p, Valor: %s\n", pointer, name)
	changeName(pointer)
	fmt.Printf("Ponteiro: %p, Valor: %s\n", pointer, name)
}
