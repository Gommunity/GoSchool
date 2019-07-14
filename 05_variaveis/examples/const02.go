package main

import (
	"fmt"
	"math"
	"time"
)

const i = "G é" + " para Go"
const j = 'V'
const k1, k2 = true, !k1
const l = 111*100000 + 9
const m1 = math.Pi / 3.141592
const m2 = 1.41421356237309504880168872420969807856967187537698078569671875376
const m3 = m2 * m2
const m4 = m3 * 1.0e+400
const n = -5.0i * 3
const o = time.Millisecond * 5

func main() {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k1, k2)
	fmt.Println(l)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	//fmt.Println(m4)
	fmt.Println(n)
	fmt.Println(o)
}
