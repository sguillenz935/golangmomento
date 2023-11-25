package main

import "fmt"

/* Declare una variable sin signo, como uint, e inicialícela con un valor negativo, por ejemplo, -10.
 Al intentar ejecutar el programa, debería obtener un error similar a este: constant -10 overflows uint. */ 

func main()  {
	
	var integer uint = -33

	fmt.Println(integer)
}