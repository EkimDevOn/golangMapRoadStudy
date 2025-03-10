package main

import ("fmt"
		"sync"
		"runtime")

func main() {

	contador := 0 // variavél em comun
	totaldegoroutines := 10

	var wg sync.WaitGroup // WaitGroup para evitar que meu programa termine antes de executar  minhas goruntines 
	wg.Add(totaldegoroutines) 

	for i := 0; i < totaldegoroutines; i++ { 
		go func() { 			  // func para disparar essa goruntines
			v := contador		 // variavél interna com o valor contador externo
		    runtime.Gosched() 	 //yield //time.Sleep(time.Second) vs. runtime.Gosched()
			v++ 				 // incrementando o valor da variavél
			contador = v 		 //salvo o valor incrementado na vaiavél
			wg.Done()      // avisa quando as funçoes forem encerradas
		}()
	}
	wg.Wait() //espera as funções terminarem antes de continuar o programa
	fmt.Println(contador) // imprime o valor da variavél contador
}
