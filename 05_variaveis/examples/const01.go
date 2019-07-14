package main

import (
	"fmt"
	"math"
	"time"
)

const a1, a2 string = "Workshop", "Go"
const b rune = 'G'
const c bool = false
const d int32 = 2019
const e float32 = 2.019
const f float64 = math.Pi * 2.0e+3
const g complex64 = 5.0i
const h time.Duration = 4 * time.Second

func main() {
	fmt.Println(a1, a2)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
