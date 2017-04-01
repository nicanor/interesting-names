package names

import "testing"
import "fmt"

func BenchmarkColor(b *testing.B) {
    for i := 0; i < b.N; i++ {
        color(15)
    }
}

func BenchmarkAnimal(b *testing.B) {
    for i := 0; i < b.N; i++ {
        color(15)
    }
}

func ExampleColor() {
    fmt.Println(color(3))
    // Output: blue
}

func ExampleAnimal() {
    fmt.Println(animal(3))
    // Output: bee
}