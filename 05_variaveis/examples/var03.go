package main

import "fmt"

var nome, desc = "Marte", "Planeta"
var raio = 6755
var massa = 641693000000000.0
var ativo = true
var satelites = []string{
	"Fobos",
	"Deimos",
}

func main() {
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", raio)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Satelites", satelites)
}
