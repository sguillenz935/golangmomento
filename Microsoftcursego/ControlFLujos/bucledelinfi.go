package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Declara una variable num de tipo int32
    var num int32

    // Obtiene la marca de tiempo actual en segundos y la asigna a la variable sec
    sec := time.Now().Unix()

    // Crea una nueva fuente de números aleatorios usando la marca de tiempo como semilla
    rand.New(rand.NewSource(sec))

    // Inicia un bucle infinito
    for {
        // Imprime un mensaje dentro del bucle
        fmt.Print("Writing inside the loop...")

        // Genera un número aleatorio entre 0 y 9 (ambos incluidos) y lo asigna a la variable num
        if num = rand.Int31n(10); num == 5 {
            // Si el número generado es 5, imprime un mensaje y sale del bucle
            fmt.Println("finish!")
            break
        }

        // Si el número no es 5, imprime el número
        fmt.Println(num)
    }
}
