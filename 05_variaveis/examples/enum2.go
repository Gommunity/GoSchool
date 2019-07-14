package main

import (
	"fmt"
)

const (
	EstrelaHiperGigante = iota
	EstrelaSuperGigante
	EstrelaBrilhanteGigante
	EstrelaGigante
	EstrelaSubGigante
)
const (
	EstrelaAna = iota
	EstrelaSubAna
	EstrelaAnaBranca
	EstrelaAnaVermelha
	EstrelaAnaMarrom
)

func main() {
	fmt.Println(EstrelaHiperGigante)
	fmt.Println(EstrelaSuperGigante)
	fmt.Println(EstrelaBrilhanteGigante)
	fmt.Println(EstrelaGigante)
	fmt.Println(EstrelaSubGigante)
	fmt.Println(EstrelaAna)
	fmt.Println(EstrelaSubAna)
	fmt.Println(EstrelaAnaBranca)
	fmt.Println(EstrelaAnaVermelha)
	fmt.Println(EstrelaAnaMarrom)
}
