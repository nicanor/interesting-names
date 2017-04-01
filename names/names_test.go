package names

import "testing"
import "fmt"

func BenchmarkFirst(b *testing.B) {
    for i := 0; i < b.N; i++ {
        animal(0)
        color(0)
        awesomeness(0)
    }
}

func BenchmarkLast(b *testing.B) {
    for i := 0; i < b.N; i++ {
        animal(31)
        color(31)
        awesomeness(31)
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

func ExampleAwesomeness() {
    fmt.Println(awesomeness(5))
    // Output: charming
}
