package main

import "fmt"

var nome, desc string
var raio int32
var massa float64
var ativo bool
var satelites []string

func main() {
	nome = "Sol"
	desc = "Estrela"
	raio = 685800
	massa = 1.989E+30
	ativo = true
	satelites = []string{
		"Mercurio",
		"Venus",
		"Terra",
		"Marte",
		"Jupiter",
		"Saturno",
		"Urano",
		"Netuno",
	}
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", raio)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Ativo? ", ativo)
	fmt.Println("Satelites", satelites)
}
