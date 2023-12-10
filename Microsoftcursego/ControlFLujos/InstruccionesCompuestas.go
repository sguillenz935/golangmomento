package main

import (
	"fmt"

)

func givemeanumber() int {
	return -1
}

func main()  {

	// El num solo se queda en el if, si se intenta declarar fuera de este aparecera un error, cosa a tener en cuenta. 
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "only has 1 digit")
	} else {
		fmt.Println(num, "has multiple numbers")
	}	

	// Esto daria error. fmt.Println(num)
}