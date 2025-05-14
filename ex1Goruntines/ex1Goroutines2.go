package main

import ("fmt"
		"sync"
)
//neste exemplo modular para selecionar quantas goroutines quer usar 

var wg sync.WaitGroup

func main() {
	novasgoroutines(100) // valor 100 ou qualquer outro desejado
	wg.Wait()
}

func novasgoroutines (i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
			go func(i int){
			fmt.Println("Eu sou goroutine 1:", i)
			wg.Done()
		}(x)
	}
}