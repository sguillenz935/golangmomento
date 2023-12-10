package main

import "fmt"

func main()  {
	x := 33
	// el % te da el resto de una division, por lo que 33 tendra un resto diferente a 0.
	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}


}