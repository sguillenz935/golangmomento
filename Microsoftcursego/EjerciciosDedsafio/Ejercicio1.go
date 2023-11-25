package main

import "fmt"

// Desafio 1
/* Establezca otra variable de tipo int y use el valor de las variables integer32 o
integer64 para confirmar el tamaño natural de la variable en el sistema. Si su sistema es de 32 bits y
usa un valor superior a 2 147 483 647, obtendrá un error de desbordamiento similar al siguiente: constant 9223372036854775807 overflows int. */

func main() {
	fmt.Println("Desafio 1")

	var integer64 int = 9223372036854775808

	fmt.Println("Esto es la variable", integer64 )


}
