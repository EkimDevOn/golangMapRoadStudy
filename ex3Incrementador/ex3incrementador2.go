package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex // utilizando mutex evitando race conditional
var contador int32

const quantidadeGoroutines = 200

func main() {

	criarGoroutines(quantidadeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:", quantidadeGoroutines,
		"Total de threads:", contador)

}

func criarGoroutines(i int) { // func que gera Goroutines
	wg.Add(i)
	for j := 0; j < i; j++ { // loop
		go func() {
			//mu.Lock() // impede que outro go run time seja executada antes de 1 acabar
			//v := contador
			atomic.AddInt32(contador, 1)
			atomic.LoadInt32(&contador)
			runtime.Gosched() // yield da thread
			//v++               // incrementando o valor da variável
			//contador = v      // salvando o valor do contador
			//mu.Unlock() // libera a proximo contador
			fmt.Println(contador)
			wg.Done()
		}()
	}
}
