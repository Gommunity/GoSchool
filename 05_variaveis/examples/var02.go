package main

import "fmt"

var nome, desc string = "Terra", "Planeta"
var raio int32 = 6378
var massa float64 = 5.972E+24
var ativo bool = true
var satelites = []string{
	"Lua",
}

func main() {
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", raio)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Satelites", satelites)
}
