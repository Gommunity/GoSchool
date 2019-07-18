# 07 - Funções

Em Go, as funções são elementos de primeira classe. Uma função sempre tem um tipo e um valor (a própria definição da função) e pode, opcionalmente, ser ligado a um identificador nomeado. Como as funções podem ser usadas como dados, elas podem ser atribuídas a variáveis ou passadas como parâmetros de outras funções.

## Declaração de uma função

Declarar uma função no Go leva a forma geral mostrada a seguir. Esta forma canônica é usada para declarar funções nomeadas e anônimas.

```go
func [<identificador da função>]([<lista de argumentos>]) [(<lista de resultados>)] {
    ...
    [return] [<valor ou lista de expressões>]
}
```

- **func**: A palavra-chave `func` marca o início da declaração da função;
- **identificador da função**: Um identificador opcional pode ser usado para nomear a função para futuras referências;
- **lista de argumentos**: É obrigatório ter um conjuto de parêntesis envolvendo uma lista de argumentos opcionais;
- **lista de resultados**: Uma lista opcional para os tipos dos valores retornados pela função. Quando mais de um valor é retornado, é obrigatório ter um conjunto de parêntesis envolvendo a lista.
- **return**: A declaração `return` faz com que o fluxo de execução saia de uma função. Quando a função define valores de retorno, uma instrução `return` é necessária como a última declaração lógica da função. Caso contrário, se nenhum valor de retorno for especificado, a instrução `return` será opcional.

A forma mais comum de definição de função em Go inclui o identificador atribuído à função. Para ilustrar isso, a seguir será mostrado o código-fonte de alguns programas com definições de funções nomeadas e diferentes combinações de parâmetros e tipos de retorno.

- Uma função com um identificador nomeado (`imprimePi`). Ela não recebe parâmetros e não retorna valores. Note que quando não há retorno, a declaração `return` é opcional.

```go
...
func imprimePi() {
	fmt.Printf("imprimePi() %v\n", math.Pi)
}
...
```

- Uma função nomeada `avogrado`. Ela não recebe parâmetros, porém, retorna um valor do tipo `float64`. Note que a declaração `return` é obrigatória quando um valor de retorno é declarado como parte da assinatura da função.

```go
...
func avogadro() float64 {
	return 6.02214129e23
}
...
```

- Uma função nomeada `fib`. Ela recebe um parâmentro `n` do tipo `int` e imprime a sequência de **Fibonacci** para `n`. Novamente, não retorna nada, portanto, a declaração `retorn` é omitida. 

```go
...
func fib(n int) {
	fmt.Printf("fib(%d): [", n)
	var p0, p1 uint64 = 0, 1
	fmt.Printf("%d %d ", p0, p1)
	for i := 2; i <= n; i++ {
		p0, p1 = p1, p0+p1
		fmt.Printf("%d ", p1)
	}
	fmt.Println("]")
}
...
```
- Uma função nomeada `isPrime`. Ela recebe um parâmetro do tipo `int` e retorna um tipo `bool`. Como a função é declarada com um retorno do tipo `bool`, a declaração do `return` é obrigatória.

```go
...
func isPrime(n int) bool {
	lim :=
		int(math.Sqrt(float64(n)))
	for p := 2; p <= lim; p++ {
		if (n % p) == 0 {
			return false
		}
	}
	return true
}
...
```

## Função variádica
Em Go, é possivel declar uma função que recebe zero ou mais argumentos. Este tipo de função é conhecida como função variádica, e ela é declarada afixando `...` (três pontos) antes do tipo do parâmetro e ele deve ser o último parâmetro. O trecho de código a seguir demonstra como declarar esse tipo de função:

```go
...
func soma(ns ...float64) float64 {
	var s float64
	for i := 0; i < len(ns); i++ {
		s += ns[i]
	}
	return s
}
...
```
O compilador resolve um parâmetro variádico como um slice.

Para invocar a função, simplesmente informe uma lista separada por virgula, como mostrado a seguir:
```go
...
func main() {
	fmt.Printf("soma = %.2f\n", soma(9, 4, 3.7, 7.1, 7.9, 9.2, 10))
}
...
```
Também é possível passar um slice para a função, neste caso, Go fornece uma sintaxe fácil para lidar com esse caso. Vamos examinar a chamada para a função `soma` no seguinte trecho de código:
```go
...
func main() {
    pontos := []]float64{9, 4, 3.7, 7.1, 7.9, 9.2, 10}
	fmt.Printf("soma = %.2f\n", pontos...)
}
...
```
Como pode ser observado, o slice pode ser passadado como um parâmetro variádico adicionando-se `...` ao parâmetro na chamada de função `suma(pontos...)`.

## Múltiplos retornos
Funções em Go podem retornar um ou mais valores. Para demonstrar esse conceito, vamos examinar o trecho de código a seguir que define uma função que implementa o algoritmo de divisão Euclidiana (https://en.wikipedia.org/wiki/Division_algorithm).

```go
...
func div(op0, op1 int) (int, int) {
	r := op0
	q := 0
	for r >= op1 {
		q++
		r = r - op1
	}
	return q, r
}
...
```
A palavra-chave `return` é seguida pela quantidade de valores correspondentes (respectivamente) aos declarados na assinatura da função. A assinatura da função `div` especifica dois valores `int` a serem retornados como resultado. Internamente, a função define as variáveis int `p` e `r` que são retornados como resultado após a conclusão da função. Esses valores retornados devem corresponder aos tipos definidos na assinatura da função ou há risco de erros de compilação.
Funções com múltiplos retornos devem ser invocadas no contexto apropriado:
- Elas devem ser atribuídas a uma lista de identificadores dos mesmos tipos, respectivamente.
- Elas só podem ser incluídas em expressões que esperam o mesmo número de retornos.

### Parâmetros de resultado nomeados
Geralmente, a lista de resultados da assinatura de uma função pode ser especificada usando identificadores nomeados junto com seus tipos. Ao usar identificadores nomeados, eles são passados para a função como variáveis regulares e podem ser acessadas e modificadas conforme necessário. Ao encontrar uma declaração `return`, os últimos valores atribuídos são retornados. Isso é demonstrado no trecho de código a seguir:
```go
...
func div(op0, op1 int) (q int, r int) {
	r = op0
	for r >= op1 {
		q++
		r = r - op1
	}
	return
}
...
```
Observe que a declaração `return` está sozinha, é omitido todos os identificadores. Para legibilidade, consistência ou estilo, você pode optar por não usar uma instrução `return` sozinha. É perfeitamente legal anexar o nome do identificador à declaração (como `return q, r`) como antes.

## Closures

## Higher-Order functions