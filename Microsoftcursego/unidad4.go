package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {

	number1, _ := strconv.Atoi(os.Args[1]) 
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("sum:", number1+number2) 
	
}

// Funciones Personalizadas

