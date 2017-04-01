package main

import "github.com/nicanor/interesting-names/names"
import "fmt"
import "math/rand"
import "time"

func main() {
    // Inicializo +rand+ con el tiempo actual, para
    // que no retorne siempre el mismo resultado.
    rand.Seed(time.Now().Unix())

    fmt.Printf(names.Randomize() + "\n")
}
