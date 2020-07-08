package main

import "fmt"



func main() {
	//funcion anonima
	a := func() {
		fmt.Println("hello!")
	}

	a()

	//otra forma de declarar la funcion anonima
	var b func()
	b = func() {
		fmt.Println("Cioa!")
	}

	b()

	// La key word "defer" difiere la ejecución de una función hasta que termine el scope del programa
	// Guarda las funciones diferidas en una pila, es decir, que se ejecutarán desde la última definida a la primera.

	defer Saludar()
	fmt.Println("Hello world!")



}

func Saludar() {
	fmt.Println("Hola mundo!")
}

	// Estructuras: en Go no existen las clases u objetos, pero las estructuras son una figura parecida. Es una colección de tipos primitivos o de otras estructuras. Posee la key word "type".
	// El zero de una estructura es el zero de todos sus componentes.
	// El nombre de la estructura o de sus componentes puede empezar con mayúscula o con minúscula. Si está en mayúscula significa que es exportable, público. De la otra forma, es privado.

	type Circle struct {
		X, Y, int
		Radius float64
		area float64
	}

	type Figure interface { //la implamentación de las interfaces en go son implícitas
		Area() float64
	}

	func (c *Circle) Area()	float64 {
		if c.area >= 0 {
			c.area = math.Pi*math.Pow(c.Radius, 2)
		}
		return c.area
	}

	func StructureFx() {

		c := &Circle{}

		c.X = 10
		c.Y = 20
		c.Radius = 5

		//otra forma de de definirlo sería c := Circle{X: 10, Y: 20, Radius: 5}
		fmt.Printf("%#v\n", c)

		ar := Area(&c) // conviene trabajar con referencia para ser más eficiente

		fmt.Printf("%#v\n", ar)
	}

