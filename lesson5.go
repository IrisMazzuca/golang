package main

import (
	"fmt"
	"sync"
)

// CHANNELS

// Se los relaciona con una tubería, donde una go routine escribe desde un lado de la misma y otra go routine lee desde el otro extremo.

// Particularidades del channel:
// - Todas las escrituras son bloqueantes, es decir, la go routine 1 se va a quedar esperando hasta que no haya una go routine leyendo y tomando su mensaje. La primera go routine se queda stand by.
// - Todo mensaje que se escribe solo se puede leer una sola vez.
// - Se escribe una vez y se lee una vez. Esto permite poder orquestar mejor.
// - Si una go routine intenta escribir en un channel cerrado --> resultado: panic.
// - No se puede reabrir un channel una vez cerrado.
// - La inicialización (creación )de un channel se realiza con la palabra reservada "make" y entre los parentesis indicamos 'chan' y luego el tipo.

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Como segundo parámetro podemos indicar la cantidad (size) de elementos que podemos permitir dentro del channel, a esto se lo conoca como 'buffer'
	messages := make(chan string, 2)

	go agent(messages, 0, &wg)
	go agent(messages, 1, &wg)

	//con el operador '<-', se indica el ingreso de información en el channel

	messages <- "Hola"
	messages <- "Hola 2"
	close(messages) // De esta forma se cierra el channel

	wg.Wait()
	//time.Sleep(100 * time.Microsecond) // Le damos tiempo para que se ejecuten las goroutines

	// messages <- "Hola 3" Escribir luego de cerrado un channel (línea 31), genera como resultado un error: Panic.

	fmt.Println("Avanza")
	//así se indica que se está leyendo del channel:
	//el segundo valor, en este caso 'open', es un booleano, indica si el channel se encuentra abierto o cerrado. (true -> abierto / false -> cerrado)

}

func agent(messages chan string, i int, wg *sync.WaitGroup) {

	defer wg.Done()
	for message := range messages {
		fmt.Println(message)
	}
	fmt.Println("Finalizó vuelta")

	// for {
	// 	message, open := <-messages
	// 	if !open {
	// 		fmt.Println("Cerró agente")
	// 		return
	// 	}
	// 	fmt.Println(message)

	// }

}
