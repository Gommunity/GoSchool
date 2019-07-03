# 05 - Variáveis e Constantes

Go é uma linguagem fortemente tipada, o que implica que todas as variáveis são elementos nomeados que estão vinculados a um valor e um tipo. Como será visto, devido a simplicidade e a flexibilidade da sintaxe da linguagem, declarar e inicializar variáveis fazem com que Go pareça mais uma linguagem tipada dinamicamente.

Antes de poder usar uma variável em Go, ela deve ser declarada com um identificador nomeado e assim poder ser referênciada futuramente no código. A forma longa para declarar uma variável em Go segue o seguinte formato:

```go
var <lista de identificadores> <tipo>
```

A palavra-chave `var` é usada para declarar um ou mais identificadores de variáveis, seguidos do seus respectivos tipos. O trecho de código a seguir mostra a declaração de diversas variáveis:

```go
...
var nome, desc string
var raio int32
var massa float64
var ativo bool
var satelites []string
...
```

## O valor zero

O trecho de código anterior mostra vários exemplos de variáveis sendo declaradas com uma variedade de tipos. À primeira vista, parece que essas variáveis não têm um valor atribuído. Na verdade, isso contradiz nossa afirmação anterior de que todas as variáveis em Go estão vinculadas a um tipo e um valor.

Durante a declaração de uma variável, se um valor não for fornecido, o compilador do Go vinculará automaticamente um valor padrão (ou um valor zero) à variável para a inicialização adequada da memória.

A tabela a seguir mostra os tipos do Go e seus valores zero padrão:

Tipo | Valor zero
-----|-----
string| "" (string vazia)
Numérico – Inteiro: byte, int, int8, int16, int32, int64, rune, uint, uint8, uint16, uint32, uint64, uintptr| 0
Numérico – Ponto flutuante: float32, float64| 0.0
booleano| false
Array| Cada índice terá um valor zero correspondente ao tipo do array.
Struct| Em uma estrutura vazia, cada membro terá seu respectivo valor zero.
Outros tipos: Interface, função, canais, slice, mapas e ponteiros| nil

## Declaração inicializada

Go também suporta a combinação de declaração de variável e inicialização como uma expressão usando o seguinte formato:

```go
var <lista de identificadores> <tipo> = <lista de valores ou expressões de inicialização>
```

Este formato de declaração possui as seguintes propriedades:

- Uma lista de identificadores fornecida no lado esquerdo do sinal de igual (seguido por um tipo)
- Uma lista de valores separados por vírgula correspondente no lado direito
- A atribuição ocorre na respectiva ordem dos identificadores e valores
- Expressões de inicialização devem produzir uma lista de valores correspondentes

O seguinte trecho de código mostra a combinação de declaração e inicialização:

```go
...
var nome, desc string = "Terra", "Planeta"
var raio int32 = 6378
var massa float64 = 5.972E+24
var ativo bool = true
var satelites = []string{
    "Lua",
}
...
```

## Omitindo o tipo das variáveis

Até agora, discutimos o que é chamado de forma longa de declaração e inicialização de variáveis em Go. Para tornar a linguagem mais próxima de linguagens tipadas dinamicamente, a especificação do tipo pode ser omitida, conforme mostrado no seguinte formato de declaração:

```go
var <lista de identificadores> = <lista de valores ou expressões de inicialização>
```

Durante a compilação, o compilador infere o tipo da variável com base no valor ou na expressão de inicialização do lado direito do sinal de igual, conforme mostrado no trecho de código a seguir.

```go
...
var nome, desc = "Marte", "Planeta"
var raio = 6755
var massa = 641693000000000.0
var ativo = true
var satelites = []string{
    "Fobos",
    "Deimos",
}
...
```

Como dito anteriormente, quando uma variável recebe um valor, ela deve receber um tipo junto com esse valor. Quando o tipo da variável é omitido, as informações de tipo são deduzidas do valor atribuído ou do valor retornado de uma expressão.

A tabela a seguir mostra o tipo que é inferido dado um valor literal:

Valor Literal | Tipo inferido
--------------|--------------
Texto com aspas duplas ou simples:<br>`"Planeta Marte"`<br><code>\`Todos os planetas giram em<br>torno do sol.\`</code> | string
Inteiros:<br>`-51`<br>`0`<br>`1234` | int
Decimal:<br>`-0.12`<br>`1.0`<br>`1.3e5`<br>`5e-11` | float64
Números complexos:<br>`-1.0i`<br>`2i`<br>`(0+2i)` | complex128
Boleanos:<br>`true`<br>`false` | bool
Arrays:<br>`[2]int{-3, 51}`| O tipo do `array` definido pelo valor literal. Neste caso `[2]int`
Map:<br>`map[string]int{`<br>`"Sol": 685800,`<br>`"Terra": 6378,`<br>`"Marte": 3396,`<br>`}` | O tipo do `map` definido pelo valor literal. Neste caso `map[string]int`
Slice:<br>`[]int{-3, 51, 134, 0}` | O tipo do `slice` definido pelo valor literal: `[]int`
Struct:<br>`struct{`<br>`nome string`<br>`diametro int`<br>`}{`<br>`"Marte", 2296,`<br>`}` | O tipo do `struct` definido conforme o valor literal. Neste caso: `struct{nome string; diametro int}`
Function:<br>`var sqr = func (v int) int {`<br>`    return v * v`<br>`}` | O tipo de `function` definido na definição literal da função. Neste caso, a variável `sqr` terá o tipo: `func (v int) int`

## Declaração curta de variável

Em Go é possível reduzir ainda mais a sintaxe da declaração de variáveis. Neste caso, usando o formato _short variable declaration_. Nesse formato, a declaração perde a palavra-chave `var` e a especificação de tipo e passa a usar o operador `:=` (dois-pontos-igual), conforme mostrado a seguir:

```go
<lista de identificadores> := <lista de valores ou expressões de inicialização>
```

Esta é uma forma simples e organizada e que é comumente usada ao declarar variáveis em Go. O techo de código a seguir mostra como usá-la:

```go
...
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
...
```
A declaração curta de variáveis utiliza o mesmo mecanismo para inferir o tipo da variável mostrado anteriormente.

### Restrições na declaração curta de variáveis

Existem algumas restrições quando usamos a declaração curta de variáveis e é muito importante estar ciente para evitar confusão:

- Em primeiro lugar, ela só pode ser usada dentro de um bloco de funções;
- o operador `:=` declara a variável e atribui os valores;
- `:=` não pode ser usado para atualizar uma variável declarada anteriormente;
- Atualizações de variáveis devem ser feitas com um sinal de igual.

## Declaração de variável em bloco

A sintaxe do Go permite que a declaração de variáveis seja agrupada em blocos para maior legibilidade e organização do código. O trecho de código a seguir mostra a reescrita de um dos exemplos anteriores usando a declaração de variável em bloco:

```go
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
```

