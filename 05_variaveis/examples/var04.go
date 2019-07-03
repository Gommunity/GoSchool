package main

import "fmt"

func main() {
	nome := "Netuno"
	desc := "Planeta"
	raio := 24764
	massa := 1.024e26
	ativo := true
	satelites := []string{
		"Náiade", "Talassa", "Despina",
		"Galateia", "Larissa", "Hipocampo", "Proteu",
		"Tritão", "Nereida", "Halimede", "Sao",
		"Laomedeia", "Neso", "Psámata",
	}
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", raio)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Ativo? ", ativo)
	fmt.Println("Satelites", satelites)
}
