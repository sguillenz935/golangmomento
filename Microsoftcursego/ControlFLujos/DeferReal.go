package main

import (
    "io"
    "os"
    "fmt"
)

func main() {
    // Crea un nuevo archivo llamado "learnGo2.txt"
    newfile, err := os.Create("learnGo2.txt")
    if err != nil {
        // Si hay un error al crear el archivo, imprime un mensaje de error y sale de la función
        fmt.Println("Error: no se pudo crear el archivo")
        return
    }

    // Se asegura de que el archivo se cierre al finalizar la función
    defer newfile.Close()

    // Escribe la cadena "learning GO" en el archivo
    if _, err := io.WriteString(newfile, "learning GO"); err != nil {
        // Si hay un error al escribir en el archivo, imprime un mensaje de error y sale de la función
        fmt.Println("Error: no se pudo escribir en el archivo")
        return
    }

    // Sincroniza los cambios en el sistema de archivos
    newfile.Sync()
}
