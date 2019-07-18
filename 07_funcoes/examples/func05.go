package main

import (
	"fmt"
)

func main() {
	fmt.Printf("soma = %.2f\n", soma(9, 4, 3.7, 7.1, 7.9, 9.2, 10))
}

func soma(ns ...float64) float64 {
	var s float64
	for i := 0; i < len(ns); i++ {
		s += ns[i]
	}
	return s
}
