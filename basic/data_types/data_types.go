/*
--- Tipos Inteiros ---
int8   : -128 to 127
int16  : -32768 to 32767
int32  : -2147483648 to 2147483647
int64  : -9223372036854775808 to 9223372036854775807

uint8  : 0 to 255
uint16 : 0 to 65535
uint32 : 0 to 4294967295
uint64 : 0 to 18446744073709551615

--- Tipos de Ponto Flutuantes ---
float32 : -3.4e+38 to 3.4e+38
float64 : -1.7e+308 to +1.7e+308

--- Tipo Booleano ---
bool    : true ou false

--- Tipos Complexos ---
complex64  : 1.401298464324817070923729583289916131280e-45 a 3.40282346638528859811704183484516925440e+38
complex128 : 4.940656458412465441765687928682213723651e-324 a 1.797693134862315708145274237317043567981e+308

--- Tipo Byte ---
byte : 0 to 255

--- Tipo Rune ---
rune : 0 to 1,114,111
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var year = 2024

	// Usando reflect
	fmt.Printf("Value: %v Type: %s \n", year, reflect.TypeOf(year))

	// Usando %t
	fmt.Printf("Value: %v Type: %T \n", year, year)
}
