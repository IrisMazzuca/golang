package main

// goroutine

import (
	"fmt"
	"sync"
)

//paquete "sync": nos permite esperar la ejecución de las goroutine sin que se termine el programa.

func main() {

	var wg sync.WaitGroup //espera que una colección de goroutines finalice. Permite de manera síncrona, agrupar ejecuciones y esperar que ese grupo de ejecuciones termine.
	wg.Add(5)             //Se deben especificar cuántos elementos hay en el grupo. En este caso 5, en el For.

	// go Imprimir() //agregando la palabra "go" delante de la función, establecemos que se trata de una goroutine

	fmt.Println("hola!")

	//se puede llamar a una goroutine con una función anónima:
	/* go func() {
		fmt.Println("Hola desde goroutine fx anónima")
	}() */

	for i := 0; i < 5; i++ {
		go func(i int) { //pasamos i con el type!
			defer wg.Done() // da aviso de la finalización de una rutina.

			fmt.Println("Hola iteración", i) // cuando se imprime, muestra todos los i = 5, esto se soliciona pasando i como parámetro e imprimirndola
		}(i)
	}

	wg.Wait() //espera a que termine la ejecución de

}

func Imprimir() {
	fmt.Println("Hola desde goroutine!")
}
