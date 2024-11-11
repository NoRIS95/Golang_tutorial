package main

import (
	"fmt"
	"math"
)

func M(p, v float64) float64 {
	return p * v
}

func W(k, p, v float64) float64 {
	m := M(p, v)
	w := math.Sqrt(k / m)
	return w
}

func T(k, p, v float64) float64 {
	w := W(k, p, v)
	return 6 / w
}

func main() {
	var k, p, v float64
	fmt.Scan(&k, &p, &v)
	t := T(k, p, v)
	fmt.Println(t)

}
