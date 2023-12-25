package main

import (
	"fmt"
)

/* En primer lugar, escriba un programa que imprima los números del 1 al 100, con los cambios siguientes:

Imprimir Fizz si el número es divisible por 3.
Imprimir Buzz si el número es divisible por 5.
Imprimir FizzBuzz si el número es divisible por 3 y por 5.
Imprimir el número si ninguno de los casos anteriores coincide.
Pruebe a usar la instrucción switch.
*/

func main() {
	fmt.Println("Ejercicio FizzBuzz")

	for i := 1; i < 101; i++ {

		switch {

		case i%5 == 0 && i%3 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}

}

// if i%3 == 0 {
// 	fmt.Println("Fizz")
// } else if i%5 == 0 {
// 	fmt.Println("Buzz")
// } else if i%5 == 0 && i%3 == 0 {
// 	fmt.Println("FizzBuzz")
// } else {
// 	fmt.Println(i)
// }
