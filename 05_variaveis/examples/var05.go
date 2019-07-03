package main

import "fmt"

var (
	nome      string  = "Terra"
	desc      string  = "Planeta"
	raio      int32   = 6378
	massa     float64 = 5.972E+24
	ativo     bool    = true
	satelites         = []string{
		"Lua",
	}
)

func main() {
	fmt.Println(nome)
	fmt.Println(desc)
	fmt.Println("Raio (km)", raio)
	fmt.Println("Massa (kg)", massa)
	fmt.Println("Satelites", satelites)
}
