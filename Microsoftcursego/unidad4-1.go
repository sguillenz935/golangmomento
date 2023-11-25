package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum", sum)

	sum22, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum22:", sum22)
	fmt.Println("Mul:", mul)

	sum33, _ := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum", sum33)

	// uso de punteros
	FirstName := "Santiagos"
	updateName(&FirstName)
	fmt.Println(FirstName)

}

// Funciones Personalizadas

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	return int1 + int2

}

// Esta es otra manera donde indicamos el resultado aunque cuidado con el return
func sum2(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}

func calc(number11 string, number22 string) (sum22 int, mul int)  {
	int11, _ := strconv.Atoi(number11)
	int22, _ := strconv.Atoi(number22)
	sum22 = int11 + int22
	mul = int11 * int22	
	return
}

// uso de punteros para actualizar el nombre gracias a una funcion
func updateName(name *string)  {
	*name = "Flagpe"
}

