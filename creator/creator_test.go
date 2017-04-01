package creator
import "testing"
import "fmt"

func BenchmarkColor(b *testing.B) {
    for i := 0; i < b.N; i++ {
        color(1)
    }
}

func BenchmarkColor2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        color2(1)
    }
}

func ExampleColor() {
    fmt.Println(color(3))
    // Output: red
}
