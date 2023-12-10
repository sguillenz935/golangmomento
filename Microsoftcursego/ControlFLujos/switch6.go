package main

import (
    "fmt"
)

// este switch tiene el falltrough, el cual lleva a la logica a pasar a la siguiente instruccion SINM VALIDAR el caso.
func main() {
    switch num := 15; {
    case num < 50:
        fmt.Printf("%d is less than 50\n", num)
        fallthrough
    case num > 100:
        fmt.Printf("%d is greater than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is less than 200", num)
    }
}