package main

import (
	"fmt"
	"math"
)

func calcHypotenuse(a, b float64) float64 {
	var hypotenuse float64 = math.Sqrt(a*a + b*b)
	return hypotenuse
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(calcHypotenuse(a, b))
}
