# GOPATH, Configuração do projeto e ativação do Go Modules

Embora seja possível criar o diretório do projeto em qualquer lugar no computador, se você planeja usar módulos para gerenciar dependências, é recomendado criar o diretório do projeto fora do diretório `GOPATH`.

Esta é uma grande mudança em relação as versões anteriores do Go (pré-1.11), onde a prática recomendada era criar o diretório dos projetos dentro de uma pasta `src` sob o diretório `GOPATH`, conforme mostrado a seguir:

```bash
$GOPATH
├── bin
├── pkg
└── src
    └── github.com
        └── <usuário github>
            └── <projeto>
```

Nessa estrutura, os diretórios possuem as seguintes funções:
- `bin`: Guardar os executáveis de nossos programas;
- `pkg`: Guardar nossas bibliotecas e bibliotecas de terceiros;
- `src`: Guardar todo o código dos nossos projetos.

## Ativação do Go Modules

Como vamos utilizar módulos, abra seu terminal e crie um novo diretório para o projeto chamado `workshop` em qualquer lugar em seu computador, desde que não esteja sob o diretório `GOPATH`.

> Dica: Crie o diretório do projeto em `$HOME/code`, mas você pode escolher um local diferente, se desejar.

```bash
$ mkdir -p $HOME/code/workshop
```

A próxima coisa que precisamos fazer é informar ao Go que queremos usar a nova funcionalidade de módulos para ajudar a gerenciar (e controlar a versão) quaisquer pacotes de terceiros que o nosso projeto importe.

Para fazer isso, primeiro precisamos decidir qual deve ser o caminho do módulo para o nosso projeto.

O importante aqui é a singularidade. Para evitar possíveis conflitos de importação com os pacotes de outras pessoas ou com a biblioteca padrão no futuro, escolha um caminho de módulo que seja globalmente exclusivo e improvável de ser usado por qualquer outra coisa. Na comunidade Go, uma convenção comum é criar o caminho do módulo com base em uma URL que você possui.

Supondo que nossa URL seja `goschool.org`, vamos iniciar os módulos da seguinte forma:

```bash
$ cd $HOME/code/workshop
$ go mod init goschool.org/workshop
go: creating new go.mod: module goschool.org/workshop
```

Neste ponto, o diretório do projeto deve possuir o aquivo `go.mod` criado.

Não há muita coisa nesse arquivo e se você abri-lo em seu editor de texto, ele deve ficar assim (mas de preferência com seu próprio caminho de módulo exclusivo):

```bash
module goschool.org/workshop

go 1.12
```

Posteriormente, veremos como esse arquivo é usado para definir os pacotes de terceiros (e suas versões).

## Informação Adicional

Se você estiver criando um pacote ou aplicativo que possa ser baixado e usado por outras pessoas e programas, é recomendável que o caminho do módulo seja igual ao local do qual o código pode ser baixado.

Por exemplo, se o seu pacote estiver hospedado em https://github.com/foo/bar, o caminho do módulo para o projeto deverá ser github.com/foo/bar.
