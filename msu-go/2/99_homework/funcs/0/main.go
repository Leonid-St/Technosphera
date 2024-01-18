package main

import (
	"fmt"
	"math"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
}
