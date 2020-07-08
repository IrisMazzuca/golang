//Ejercicio Clase 3
// 1. Producto (Nombre , Precio)
// 2. Carrito (cantidad , Producto)
// 3. Calcular total
package main

import "fmt"

type Producto interface {
	Nombre() string
	Precio() float32
}

func (l *Leche) Nombre() string {
	return l.nombre
}

func (l *Leche) Precio() float32 {
	return l.precio
}

func (c *Cafe) Nombre() string {
	return l.nombre
}

func (c *Cafe) Precio() float32 {
	return l.precio
}

func main() {
	l := &Leche{Nombre: "Sancor", Precio: 80}
	c := &Cafe{Nombre: "La Morenita", Precio: 300}

	carrito := []*Item{}

	carrito = append(carrito, &Item {
		Cantidad: 2,
		Producto: l,
	})

	carrito = append(carrito, &Item {
		Cantidad: 1,
		Producto: c,
	})
	

	precioTotal := 0
	for _, item := range carrito {
		precioTotal += item.Producto.Precio() * item.Cantidad
	}

	fmt.Println("Precio total:", precioTotal)

}