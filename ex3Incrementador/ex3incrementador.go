package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Variável compartilhada para o contador
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Número de goroutines
	numGoroutines := 1000

	// Adiciona as goroutines ao WaitGroup
	wg.Add(numGoroutines)

	// Cria e inicia as goroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			mu.Lock()
			defer wg.Done() // Indica que a goroutine terminou

			// Lê o valor do contador
			val := counter

			// Faz yield da thread
			runtime.Gosched()

			// Incrementa o valor salvo
			val++

			// Copia o novo valor para a variável inicial
			counter = val
			mu.Unlock()
		}()
	}

	// Aguarda todas as goroutines terminarem
	wg.Wait()

	// Imprime o valor final do contador
	fmt.Printf("Valor final do contador: %d\n", counter)
}
