package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Inicializa um WaitGroup para sincronizar as goroutines
	var wg sync.WaitGroup

	// Adiciona 2 ao WaitGroup, uma para cada goroutine adicional
	wg.Add(2)

	// Cria a primeira goroutine adicional
	go func() {
		defer wg.Done() // Indica que a goroutine terminou
		fmt.Println("Goroutine 1: Iniciando...")
		time.Sleep(1 * time.Second) // Simula uma tarefa que leva tempo
		fmt.Println("Goroutine 1: Terminando...")
	}()

	// Cria a segunda goroutine adicional
	go func() {
		defer wg.Done() // Indica que a goroutine terminou
		fmt.Println("Goroutine 2: Iniciando...")
		time.Sleep(1 * time.Second) // Simula uma tarefa que leva tempo
		fmt.Println("Goroutine 2: Terminando...")
	}()

	// Aguarda até que todas as goroutines adicionais terminem
	wg.Wait()

	// Apenas a goroutine principal imprime esta mensagem
	fmt.Println("Programa principal: Todas as goroutines terminaram.")
}
