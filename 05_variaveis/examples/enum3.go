package main

import (
	"fmt"
)

const (
	EstrelaAna byte = iota
	EstrelaSubAna
	EstrelaAnaBranca
	EstrelaAnaVermelha
	EstrelaAnaMarrom
)

func main() {
	fmt.Println(EstrelaAna)
	fmt.Println(EstrelaSubAna)
	fmt.Println(EstrelaAnaBranca)
	fmt.Println(EstrelaAnaVermelha)
	fmt.Println(EstrelaAnaMarrom)
}
