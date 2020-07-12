package main

import (
	"fmt"
	"sync"
)

// Mutex - zona de exclusión mutua
// Se usan en programación concurrente para evitar que entre más de un proceso a la vez en la sección crítica. La sección crítica es el fragmento de código donde puede modificarse un recurso compartido.

// mu.Lock() -->
// acceder a resultados
// zona de exclusión mutua
// evitar los fatal errors y cuidar el espacio en la memoria
// mu.Unlock() --> desbloqueo

func main() {
	resultados := map[string]float64{}

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			mu.Lock()
			resultados[fmt.Sprintf("resultado - %d", i)] = float64(i)
		}(i)
	}
	wg.Wait()
}
