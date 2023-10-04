package main

import "fmt"

func main() {

	const pi float64 = 3.14159265

	const pi2= 3.14159265

	fmt.Printf("%.6f \n", pi)
	fmt.Printf("%.4f \n", pi2)

	fmt.Println("Hello World")


	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)
}