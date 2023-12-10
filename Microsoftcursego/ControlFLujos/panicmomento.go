package main

import "fmt"

func highlow(high int, low int) {
    // Verifica si el valor de 'high' es menor que 'low'
    if high < low {
        fmt.Println("Panic!")
        // Lanza una panic con un mensaje y termina la ejecución del programa
        panic("highlow() low greater than high")
    }

    // Deferred statement que se ejecutará al finalizar la función
    defer fmt.Println("Deferred: highlow(", high, ",", low, ")")

    // Imprime un mensaje antes de realizar la llamada recursiva
    fmt.Println("Call: highlow(", high, ",", low, ")")

    // Llamada recursiva a la función highlow con 'low' incrementado en 1
    highlow(high, low+1)
}

func main() {
    // Llamada inicial a la función highlow con high=2 y low=0
    highlow(2, 0)

    // Mensaje impreso después de que highlow() termina debido a la recursión infinita
    fmt.Println("Program finished successfully!")
}
