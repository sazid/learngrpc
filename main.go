package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberItems[T any](n int, f func() T) []T {
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = f()
	}
	return res
}

func NewInt() int {
	return rand.Intn(10)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println(randomNumberItems(rand.Intn(10), NewInt))
}
