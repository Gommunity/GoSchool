# 05 - Variáveis, Constantes e Enumerações

## Variáveis

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

### O valor zero

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

### Declaração inicializada

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

### Omitindo o tipo das variáveis

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

### Declaração curta de variável

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

#### Restrições na declaração curta de variáveis

Existem algumas restrições quando usamos a declaração curta de variáveis e é muito importante estar ciente para evitar confusão:

- Em primeiro lugar, ela só pode ser usada dentro de um bloco de funções;
- o operador `:=` declara a variável e atribui os valores;
- `:=` não pode ser usado para atualizar uma variável declarada anteriormente;
- Atualizações de variáveis devem ser feitas com um sinal de igual.

### Declaração de variável em bloco

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

## Constantes

Em Go, uma constante é um valor com uma representação literal de uma string, um caractere, um boleano ou numeros. O valor para uma constante é estático e não pode ser alterado após a atribuição inicial.

### Constantes tipadas

Go usa a palavra chave `const` para indicar a declaração de uma constante. Diferente da declaração de uma variável, a declaração deve sempre incluir o valor literal a ser vinculado ao identificador, conforme mostrado a seguir:

```go
const <lista de identificadores> tipo = <lista de valores ou expressões de inicialização>
```

O seguinte trecho de código mostra algumas constantes tipadas sendo declaradas:

```go
...
const a1, a2 string = "Workshop", "Go"
const b rune = 'G'
const c bool = false
const d int32 = 2019
const e float32 = 2.019
const f float64 = math.Pi * 2.0e+3
const g complex64 = 5.0i
const h time.Duration = 4 * time.Second
...
```

Note que cada constante declarada recebe explicitamente um tipo. Isso implica que a constantes só podem ser usada em contextos compatíveis com seus tipos. No entanto, isso funciona de maneira diferente quando o tipo é omitido.

### Constantes não tipadas

Constantes são ainda mais interessantes quando não são tipada. Uma constante sem tipo é declarada da seguinte maneira:

```go
const <lista de identificadores> = <lista de valores ou expressões de inicialização>
```

Neste formato, a especificação de tipo é omitida na declaração. Logo, uma constante é meramente um bloco de bytes na memória sem qualquer tipo de restrição de precisão imposta. A seguir, algumas declarações de constantes não tipificadas:

```go
...
const i = "G é" + " para Go"
const j = 'V'
const k1, k2 = true, !k1
const l = 111*100000 + 9
const m1 = math.Pi / 3.141592
const m2 = 1.41421356237309504880168872420969807856967187537698078569671875376
const m3 = m2 * m2
const m4 = m3 * 1.0e+400
const n = -5.0i * 3
const o = time.Millisecond * 5
...
```

A constante `m4` recebe um valor muito grande (`m3 * 1.0e+400`) que é armazenado na memória sem qualquer perda de precisão. Isso pode ser útil em aplicações onde realizar cálculos com um alto nível de precisão é extremamente importante.

### Atribuindo constantes não tipadas

Mesmo Go sendo uma linguagem fortemente tipada, é possível atribuir uma constante não tipada a diferentes tipos de precisão diferentes, embora compatíveis, sem qualquer reclamação do compilador, conforme mostrada a seguir:

```go
...
const m2 = 1.41421356237309504880168872420969807856967187537698078569671875376
var u1 float32 = m2
var u2 float64 = m2
u3 := m2
...
```

O exemplo anterior mostra a constante não tipada `m2` sendo atribuída a duas variáveis de ponto flutuante com diferentes precisões, `u1` e `u2`, e a uma variável sem tipo, `u3`. Isso é possível porque a constante `m2` é armazenada como um valor não tipado e, portanto, pode ser atribuída a qualquer variável compatível com sua representação (um ponto flutuante).

Como `u3` não tem um tipo específico, ele será inferido a partir do valor da constante, e como `m2` representa um valor decimal, o compilador irá inferir seu tipo padrão, um `float64`.

A declaração de constantes também podem ser organizadas em blocos, aumentando a legibilidade do código, conforme a seguir:

```go
...
const (
	a1, a2 string        = "Workshop", "Go"
	b      rune          = 'G'
	c      bool          = false
	d      int32         = 2019
	e      float32       = 2.019
	f      float64       = math.Pi * 2.0e+3
	g      complex64     = 5.0i
	h      time.Duration = 4 * time.Second
)
...
```

## Enumerações

Um interessante uso para constantes é na criação de enumerações. Usando a declaração de blocos, é facilmente possível criar valore inteiros que aumentam numericamente. Para isso, basta atribuir o valor constante pré-declarado `iota` a um identificador de constante na declaração de bloco, conforme mostrado no exemplo a seguir:

```go
...
const (
	EstrelaHiperGigante = iota
	EstrelaSuperGigante
	EstrelaBrilhanteGigante
	EstrelaGigante
	EstrelaSubGigante
	EstrelaAna
	EstrelaSubAna
	EstrelaAnaBranca
	EstrelaAnaVermelha
	EstrelaAnaMarrom
)
...
```

Nessa situação, o compilador fará o seguinte:

- Declarar cada membro no bloco como um valor constante inteiro não tipado;
- Inicializar a `iota` com o valor zero;
- Atribuir a `iota`, ou zero, ao primeiro membro (`EstrelaHiperGigante`);
- Cada constante subsequente recebe um `int` aumentado em um.

Assim, as constantes da lista receberão os valores de zero até nove.

É importante ressaltar que, sempre que `const` aparecer em um bloco de declaração, o contador é redefinido para zero. No trecho de código seguinte, cada conjunto de constantes é enumerado de zero a quatro:

```go
...
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
...
```

### Substituindo o tipo padrão de uma enumeração

Por padrão, uma constante enumerada é declarada como um tipo inteiro não tipado. Porém, podemos substituir o tipo padrão provendo explicitamente um tipo numérico, como mostrado a seguir:

```go
...
const (
	EstrelaAna byte = iota
	EstrelaSubAna
	EstrelaAnaBranca
	EstrelaAnaVermelha
	EstrelaAnaMarrom
)
...
```
É possível especificar qualquer tipo numérico que pode representar um inteiro ou um ponto flutuante. No exemplo anterior, cada constante será declarada como um tipo `byte`.

### Usando `iota` em expressões

Quando a `iota` aparece em uma expressão, o compilador irá aplicar a expressão para cada valor sucessivo. O exemplo a seguir atribui números pares aos membros do bloco de declaração:

```go
...
const (
	EstrelaHiperGigante = 2.0 * iota
	EstrelaSuperGigante
	EstrelaBrilhanteGigante
	EstrelaGigante
	EstrelaSubGigante
)
...
```

### Ignorando valores em enumerações

É possível ignorar certos valores em uma enumeração simplesmente atribuindo a `iota` a um identificador em branco (`_`). No trecho de código a seguir, o valor `0` é ignorado:

```go
...
const (
	_                   = iota
	EstrelaHiperGigante = 1 << iota
	EstrelaSuperGigante
	EstrelaBrilhanteGigante
	EstrelaGigante
	EstrelaSubGigante
)
...
```