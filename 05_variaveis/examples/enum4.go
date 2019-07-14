package main

import (
	"fmt"
)

const (
	EstrelaHiperGigante = 2.0 * iota
	EstrelaSuperGigante
	EstrelaBrilhanteGigante
	EstrelaGigante
	EstrelaSubGigante
)

func main() {
	fmt.Println(EstrelaHiperGigante)
	fmt.Println(EstrelaSuperGigante)
	fmt.Println(EstrelaBrilhanteGigante)
	fmt.Println(EstrelaGigante)
	fmt.Println(EstrelaSubGigante)
}
