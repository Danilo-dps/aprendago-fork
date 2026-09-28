# Golang

## Recursos

- [Effective Go](https://go.dev/doc/effective_go)
- [Blog Go](https://go.dev/blog/all)
- [Especificação](https://go.dev/ref/spec)
- [Built-in](https://pkg.go.dev/builtin)
- [Golang doc](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)

---

## `new` e `make`

Segundo a própria documentação do Go, as built-in functions `new` e `make` são usadas para alocar memória para um tipo ([Effective Go, Allocation](https://go.dev/doc/effective_go#data)). Elas fazem coisas diferentes:

- `new(T)`: aloca memória zerada para `T` e devolve um **ponteiro** `*T`. Funciona com qualquer tipo.
- `make(T, args)`: só para **slices, maps e canais**. Inicializa a estrutura interna desses tipos e devolve o próprio `T` (não um ponteiro).

```go
p := new(int)             // *int apontando para 0
s := make([]int, 0, 10)   // slice com len 0 e cap 10
m := make(map[string]int) // map pronto para uso
c := make(chan int)       // canal bidirecional
```

`make` existe porque slices, maps e canais referenciam estruturas internas que precisam ser inicializadas. Um map com valor zero (`nil`) causa panic ao escrever; um canal `nil` bloqueia para sempre.

---

## `append`

Adiciona elementos ao final de um slice e **devolve o slice resultante**, por isso o retorno precisa ser reatribuído.

```go
s := []int{1, 2}
s = append(s, 3)        // [1 2 3]
s = append(s, 4, 5)     // [1 2 3 4 5]
s = append(s, outro...) // concatena outro slice
```

- Se a capacidade (`cap`) é insuficiente, o Go aloca um novo array subjacente e copia os elementos.
- Ao passar um slice para uma função, o `append` feito lá dentro altera apenas a cópia local (`len`/`cap`/ponteiro). Para o chamador enxergar, a função deve **retornar** o slice (`s = adiciona(s)`).

---

## `range` (sintaxe do `for` com `range`)

`for ... range` percorre diferentes tipos, e o que cada iteração devolve depende do tipo:

| Tipo | Valores da iteração |
| --- | --- |
| slice / array | índice, valor (cópia) |
| string | índice do byte, `rune` |
| map | chave, valor (ordem **aleatória**) |
| canal | valor; termina quando o canal é **fechado** |
| inteiro (Go 1.22+) | 0 até n-1 |

```go
for i, v := range slice {}
for k, v := range m {}
for i, r := range "olá" {} // i é posição em bytes, r é rune
for v := range canal {}    // só termina com close(canal)
for i := range 5 {}        // 0..4
for range 3 {}             // só repete 3 vezes, sem usar o índice
```

- Use `_` para ignorar um dos valores: `for _, v := range s`.
- `v` é uma **cópia** do elemento; alterar `v` não altera o slice.
- A partir do Go 1.22, cada iteração tem sua própria variável de loop.

---

## `rune`

`rune` é um alias de `int32` e representa um **code point Unicode**. Uma `string` em Go é uma sequência de **bytes** (UTF-8), e um caractere pode ocupar mais de um byte.

```go
s := "ação"
fmt.Println(len(s))         // 6 (bytes)
fmt.Println(len([]rune(s))) // 4 (caracteres)

var r rune = 'ç' // literal com aspas simples é rune
```

- `range` sobre string decodifica UTF-8 e entrega `rune`.
- Indexar (`s[i]`) devolve um `byte`, não um caractere.

---

## Struct embutida vs struct como campo

**Campo nomeado**: acesso pelo nome do campo.

**Embutida (embedding)**: o tipo é declarado sem nome de campo, e seus campos e métodos são **promovidos** para a struct externa. É composição, não herança.

```go
type Animal struct{ Nome string }

func (a Animal) Apresentar() string { return "Sou " + a.Nome }

// embutida
type Cachorro struct {
 Animal
 Raca string
}

// campo nomeado
type Dono struct {
 Nome string
 Pet  Animal
}

c := Cachorro{Animal: Animal{Nome: "Rex"}, Raca: "Vira-lata"}
c.Nome           // promovido
c.Apresentar()   // método promovido
c.Animal.Nome    // também funciona

d := Dono{Nome: "Ana", Pet: Animal{Nome: "Rex"}}
d.Pet.Nome       // sempre pelo campo
```

- Em conflito de nomes, vale o campo/método de menor profundidade (o da struct externa "esconde" o do embutido).
- Métodos promovidos fazem a struct externa satisfazer interfaces.

---

## Ponteiros: quando usar `&` e quando usar `*`

- `&x`: **obtém o endereço** de `x` (produz um ponteiro).
- `*T` em um **tipo**: significa "ponteiro para T".
- `*p` em uma **expressão**: **desreferencia**, lê ou escreve o valor apontado.

```go
x := 10
p := &x        // p é *int
*p = 20        // altera x através do ponteiro
fmt.Println(x) // 20

func dobra(n *int) { *n *= 2 }
dobra(&x)      // passa o endereço
```

- Para acessar campo de struct por ponteiro, o Go desreferencia automaticamente: `p.Campo` equivale a `(*p).Campo`.
- `&Pessoa{...}` cria a struct e já devolve o ponteiro.
- Go não tem aritmética de ponteiros.

---

## Receivers: valor vs ponteiro

A regra sobre ponteiros versus valores para *receivers* é que métodos que recebem valor podem ser invocados tanto em ponteiros quanto em valores, mas métodos que esperam ponteiros só podem ser invocados em ponteiros ([Effective Go](https://go.dev/doc/effective_go#pointers_vs_values)).

Há uma exceção prática: quando o valor é **endereçável** (uma variável, por exemplo), o compilador insere o `&` automaticamente.

```go
type Contador struct{ n int }

func (c Contador) Valor() int   { return c.n } // receiver valor (cópia)
func (c *Contador) Incrementa() { c.n++ }      // receiver ponteiro

v := Contador{}
v.Incrementa() // ok: v é endereçável, vira (&v).Incrementa()
p := &v
p.Valor()      // ok: vira (*p).Valor()

// Contador{}.Incrementa()  // erro: valor não endereçável
// m["a"].Incrementa()      // erro: elemento de map não é endereçável
```

- Um método com receiver ponteiro altera o original; com receiver valor, altera só uma cópia.
- Para interfaces, o conjunto de métodos de `T` inclui só os receivers valor, e o de `*T` inclui todos. Então, se `Incrementa` faz parte de uma interface, é o `*Contador` que a satisfaz.
- Convenção: se algum método precisa de receiver ponteiro, use ponteiro em todos, por consistência.

---

## Ponteiros vs valores: quando faz sentido

Ponteiros fazem mais sentido com tipos de semântica de valor (structs, arrays, tipos básicos) quando a função precisa alterar o original, quando o valor é grande, ou quando o tipo não pode ser copiado (`sync.WaitGroup`, `sync.Mutex`). Canais, maps, funções e interfaces já carregam referência, então quase nunca precisam de ponteiro. Slices compartilham os elementos, mas `append` na cópia não afeta o slice original (retornar o slice resolve).

| Tipo | O que é copiado ao passar | Compartilha dados? |
| --- | --- | --- |
| `chan T` | referência ao canal | Sim |
| `map[K]V` | referência ao map | Sim |
| `[]T` | ponteiro + len + cap | Elementos sim; `append` não |
| `*T` | endereço | Sim |
| struct, array, `int`, `string` | os dados inteiros | Não |

Go é sempre **pass by value**; o que muda é *o que* o valor contém.

---

## Filosofia de concorrência

`Do not communicate by sharing memory; instead, share memory by communicating.`

Em vez de várias goroutines acessando a mesma memória protegida por locks, prefira **enviar os dados entre goroutines por canais**, de modo que apenas uma goroutine seja "dona" do dado por vez. (Mutex continua sendo uma ferramenta válida quando um lock simples resolve melhor.)

---

## Comma ok

Idioma em que uma operação devolve um segundo valor booleano indicando se ela "deu certo".

```go
// 1. Map: a chave existe?
v, ok := m["chave"]

// 2. Type assertion: o valor é desse tipo?
s, ok := i.(string)

// 3. Canal: o valor veio de um envio real?
v, ok := <-canal // ok == false: canal fechado e vazio (v é o valor zero)
```

- Sem o `ok`, receber de um canal fechado devolve o valor zero imediatamente e para sempre, o que pode gerar valores espúrios em um `select`.
- Truque em `select`: ao detectar `ok == false`, atribuir `nil` à variável do canal desativa aquele `case` (um canal `nil` nunca fica pronto).

```go
case v, ok := <-c:
 if !ok {
  c = nil
  continue
 }
```

---

## Canais

### Direção

- O canal nasce sempre bidirecional (`make(chan T)`). A direção (`chan<- T` / `<-chan T`) é uma restrição aplicada nas assinaturas de função (ou variáveis) e checada pelo compilador, definindo o papel de cada uma: só envia ou só recebe.
- Criar um canal já restrito (`make(<-chan T)`) compila, mas é inútil: ninguém teria como enviar nele, então os receives bloqueariam para sempre.

### Conversão entre direções

Um canal bidirecional (`chan T`) pode ser atribuído/passado para um tipo direcional (`chan<- T` ou `<-chan T`), mas o contrário não é permitido. Também não há conversão entre os dois tipos direcionais. Ou seja, a direção só pode ser restringida, nunca ampliada ou trocada.

```go
c := make(chan int)

var s chan<- int = c // ok
var r <-chan int = c // ok

// c = s // erro
// r = s // erro
```

### Semântica de referência

Canais têm semântica de referência: a variável já carrega uma referência ao canal real, então passar a variável (cópia) basta para enviar, receber e fechar. Reatribuir a variável (`c = nil`) muda só a cópia local. Ponteiro para canal (`*chan T`) quase nunca faz sentido.

### Sincronização

Além de transportar dados, o canal também sincroniza as goroutines: um envio bloqueia até haver receptor, e um receive bloqueia até haver dado ou fechamento.
