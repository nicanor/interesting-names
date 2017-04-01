package main

import "github.com/nicanor/interesting-names/creator"
import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().Unix())
    fmt.Printf(creator.Randomize())
    fmt.Printf("\n%d\n", creator.GetNum())
    fmt.Printf("%d\n", creator.GetNum())
    fmt.Printf("%d\n", creator.GetNum())
    fmt.Printf("%d\n", creator.GetNum())
    fmt.Printf("%d\n", creator.GetNum())
    fmt.Printf("%d\n", creator.GetNum())
}
