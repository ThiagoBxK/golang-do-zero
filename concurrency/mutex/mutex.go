package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex // Evita que várias goroutines modifiquem a variável ao mesmo tempo, dando Lock ou Unlock
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Valor final do contador:", counter)
}
