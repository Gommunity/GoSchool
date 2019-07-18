package main

import (
	"fmt"
)

func div(op0, op1 int) (q int, r int) {
	r = op0
	for r >= op1 {
		q++
		r = r - op1
	}
	return
}

func main() {
	q, r := div(71, 5)
	fmt.Printf("div(71,5) -> q = %d, r = %d\n", q, r)
}
