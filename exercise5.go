package main

import (
	"fmt"
	"strings"
	"sync"
)

// var concurrency = flag.Int("concurrency")

func main() {
	var wg sync.WaitGroup

	// concurrency := 2

	frases := []string{
		"Chicas programando en Golang",
		"Chicas programando los jueves a las 20hs",
		"The Wildcast trasmite a través de Twitch",
		"Gophers <3",
	}

	frase := "Chicas programando"

	wg.Add(len(frases))

	messages := make(chan string)

	for elem := range frases {
		go chequearString(messages, &wg, frase)
		message <- elem

	}
	close(messages)
	wg.Wait()

}

func chequearString(messages chan string, wg *sync.WaitGroup, frase string) {
	defer wg.Done()
	for message := range messages {
		if strings.Contains(message, frase) {
			fmt.Println(message)
		}

	}
}
