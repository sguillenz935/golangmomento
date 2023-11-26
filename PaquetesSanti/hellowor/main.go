package main

import (
  "fmt"
  "github.com/myuser/calculadora"
)

func main() {
    total := calculator.Sum(3, 5)
    fmt.Println(total)
    fmt.Println("Version: ", calculator.Version)
    
}