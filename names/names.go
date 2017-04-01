package names

import "math/rand"

func Create(c, a int) string {
    return color(c) + "-" + animal(a)
}

func Randomize() string {
    return Create(rand.Intn(32), rand.Intn(32))
}

func GetNum() int {
    return rand.Intn(32)
}
