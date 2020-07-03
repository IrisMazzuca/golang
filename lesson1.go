package main

import (
	"fmt"
	"math/rand"
)

/*
Consigna:
- generar un número aleatorio
- usuario debe igresar un número a fin de adivinar
- si es más chico -> devuelve msj que es más chico, si es más grande -> devuelve msj que es más grande
- si adivina el usuario, termina el programa
- si no adivina, suma 1 a la cantidad de intentos
*/

func main() {

	numeroParaAdivinar := rand.Intn(11) //número de al azar del 0 al 10

	Jugar(numeroParaAdivinar)
}

func Jugar(numeroParaAdivinar int) {

	intentos := 0

	fmt.Println("Ingrese un número del 0 al 10")

	var numeroInputUsuario int

	for {
		fmt.Scanln(&numeroInputUsuario) //espera a que apretes enter y captura lo que ingresaste en la terminal

		if numeroInputUsuario > numeroParaAdivinar {
			fmt.Println("El número es menor")
			intentos++
		} else if numeroInputUsuario < numeroParaAdivinar {
			fmt.Println("El número es mayor")
			intentos++
		} else if numeroInputUsuario == numeroParaAdivinar {
			fmt.Println("Adivinaste! :) Cantidad de intentos hasta adivinar:")
			fmt.Println(intentos)
			break
		}
	}
}

// func main() {
// 	saludo := CrearSaludo("Iris")

// 	fmt.Println(saludo)
// }

// func CrearSaludo(nombre string) string {
// 	return fmt.Sprintln("Hola", nombre)
// }

// Ejemplo 1 FOR:

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// Ejemplo 2 FOR:

// func main() {
// 	i := 0
// 	for {
// 		fmt.Println(i)
// 		i++

// 		if i > 10 {
// 			break
// 		}
// 	}
// }
