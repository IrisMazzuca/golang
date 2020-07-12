package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type producto float64 // puedo crear un tipo "producto" que es un float64. Esto me permite crear métodos

// Precio retorna precio del producto
func (p producto) Precio() float64 {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	return float64(p) //retorna el precio
}

func main() {
	productos := map[string][]producto{
		"coto":      []producto{123.10, 97.9, 99.2, 40.3, 57},
		"dia":       []producto{92.3, 83, 99, 540, 223},
		"carrefour": []producto{554, 12, 47.5, 92.3},
	}

	// 1) Iterar productos
	// 2) Por cada supermercado iterar los precios y sumarlos
	// 3) Tenemos que esperar a que finalice cada iteración

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	resultados := map[string]float64{}
	wg.Add(3)
	for super, items := range productos {
		go func(super string, items []producto) {
			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			totalSuper := 0.0
			for _, producto := range items {
				totalSuper += producto.Precio()
			}
			mu.Lock()
			resultados[super] = totalSuper
		}(super, items)

	}
	wg.Wait()
	fmt.Printf("%+v\n", resultados)
}
