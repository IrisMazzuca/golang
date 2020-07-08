//lesson 2

package main

import "fmt"

func main() {
	var a string
	a = "hello"

	b := &a

	fmt.Println(a, b) // esto imprime en consola: hello 0xc0001041e0 (b imprime la referencia en memoria)

	fmt.Println(a, *b) // esto imprime en consola: hello hello (con el asterisco solicito el valor de referencia)

	Saludar(&a) // llamo a la fx saludar y le paso como parámetro la referencia a 'a'

	Arrays()

	Slices()
}

func Saludar(saludo *string) {
	fmt.Println(*saludo)
}

// las referencias se realizan a variables, NO a un string literal
// el beneficio de la referencia es el ahorro en memoria

// valores por defecto de los tipos:
// string: cadena vacia ""
// int: 0
// booleano: false
// en el caso de una referencia, por defecto es null

//---------

// Array -> arreglo de un tipo de dato. En golang siempre tienen un tamaño específico (no se puede aumentar o disminuir). Un array es más eficiente que un slice.

func Arrays() {
	c := [2]int{1, 2} // el dato entre corchetes indica el lenght del array, luego el tipo de dato que contiene y entre llaves los valores.

	d := [2]int{} // puedo declararlo sin valores originales
	d[0] = 12
	d[1] = 100

	fmt.Println("array 1:", c, "array 2:", d)
}

//---------

// Slice -> tiene un valor dinámico. Siempre tiene de fondo un array y me permite ver una porcion de él. Si quiero agregar un valor, se puede instanciar.

func Slices() {
	f := [4]int{1, 2, 3, 4} // array

	e := []int{25, 48} // sintaxis de slice (corchetes vacios)

	e = append(e, 90) // con append agrego valores al slice

	fmt.Println("slice:", e)

	// g := f[:] // se crea un slice con todos los valores del array f. Si se quiere especificar los valores a pasar: g := f[1:3]

	// Iteración de un array

	for i := 0; i < len(f); i++ {
		fmt.Println(f[i])
	}

	// otra forma de utilizar el for en golang

	for i, v := range f {
		fmt.Println("iteracion con range:", i, v)
	}

}

//---------

// Map

func Maps() {

	//formas de declarar un map

	var m map[string]int

	m["a"] = 30
	m["b"] = 31
	m["c"] = 32

	m2 := map[int]string{
		0: "hello",
		1: "world",
		2: " !!",
	}

	fmt.Println(m["b"])

	delete(m, "b") // elimina una key

	fmt.Printf("%#v\n", m2) // con Printf le das formato al print

	// Iteración de un map (for con key y valor)
	// los mapas en go no vienen en orden cuando lo iteras, siempre es orden aleatorio

	for k, v := range m2 {
		fmt.Println(k, v)
	}

}
