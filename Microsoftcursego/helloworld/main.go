package main

import (
  "fmt"
  "github.com/sguillenz935/calculator"
)

func main() {
    total := calculator.Sum(3, 5)
    fmt.Println(total)
    fmt.Println("Version: ", calculator.Version)
    fmt.Println(calculator.Version)
    fmt.Println(calculator.Version)
}