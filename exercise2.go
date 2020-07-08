package main

import "fmt"

// Ejercicio: se pide al usuario un número, iteramos ese num desde cero para encontrar los numeros pares y poder mostrarlos.

func main() {

	fmt.Println("Ingrese un número:")
	var numeroIngresado int
	fmt.Scanln(&numeroIngresado)
	pares := identificarPares(numeroIngresado)

	for k, v := range pares {
		fmt.Printf("%d : %d", k, v)
	}
}

func identificarPares(numeroIngresado int) []int {
	pares := []int{}

	for i := 0; i <= numeroIngresado; i++ {
		if i%2 == 0 {
			pares = append(pares, i)
		}
	}

	fmt.Printf("%+v", pares)
	return pares
}
