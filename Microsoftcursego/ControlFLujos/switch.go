package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix()
	rand.New(rand.NewSource(sec))
	i := rand.Int31n(10)

	fmt.Println("Número generado:", i)

	switch i {
	case 0:
		fmt.Println("zero...")
	case 1:
		fmt.Println("one...")
	case 2:
		fmt.Println("two...")
	default: 
		fmt.Println("No match...")
		
	}

	fmt.Println("ok")

}
