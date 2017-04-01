package creator

import "math/rand"

func Create(x int) string {
    return color(x) + "-" + animal(x)
}

func Randomize() string {
    return Create(rand.Intn(32))
}

func GetNum() int {
    return rand.Intn(32)
}
