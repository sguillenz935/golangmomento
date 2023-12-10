package main

import (
	"fmt"
    "math/rand"
    "time"
)

func main() { 
	var num int64
	// Crea una nueva fuente de números aleatorios usando la hora actual como semilla
	rand.New(rand.NewSource(time.Now().Unix()))
	for num !=5 {
		// Genera un número aleatorio entre 0 y 14 (ambos incluidos) y lo asigna a num
		num = rand.Int63n(15)
		fmt.Println(num)
	}

}