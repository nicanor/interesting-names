package names

import "math/rand"

// Get a random Colorful-Awesome-Animal
func Randomize() string {
    return Create(rand.Intn(32), rand.Intn(32), rand.Intn(32))
}

// Get a specific Colorful-Awesome-Animal
func Create(q, c, a int) string {
    return color(c) + "-" + quality(q) + "-" + animal(a)
}
