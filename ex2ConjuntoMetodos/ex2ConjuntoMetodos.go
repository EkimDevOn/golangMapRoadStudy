package main

import "fmt"

// Definindo o tipo pessoa
type pessoa struct {
	nome string
}
// O conjunto de métodos tem distinção entre reciver *ponteiro eo reciver  não-ponteiro!

// Definindo o método falar com receiver ponteiro (*pessoa)
func (p *pessoa) falar() { // sem argumentos e sem retornos
	fmt.Printf("Olá, meu nome é %s\n", p.nome)
}

// Definindo a interface humanos
type humanos interface {
	falar()
}

// Definindo a função dizerAlgumaCoisa que aceita um parâmetro do tipo humanos
func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	// Criando uma pessoa
	p1 := pessoa{nome: "João"}

	p1.falar() // É um shortcut pra (&p1).falar()

	// Utilizando um valor do tipo *pessoa na função dizerAlgumaCoisa
	dizerAlgumaCoisa(&p1) // Funciona porque &p1 é do tipo *pessoa, que implementa humanos

	(&p1).falar() // Maneira mais "correta"

	// dizerAlgumaCoisa(p1) => nao funciona!

	// Tentando utilizar um valor do tipo pessoa na função dizerAlgumaCoisa
	// dizerAlgumaCoisa(p1) // Isso causará um erro de compilação porque p1 é do tipo pessoa, não *pessoa
}
