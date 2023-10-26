package clase1

import "fmt"

func clase1() {

	const pi float32 = 3.14159265

	const pi2= 3.14159265

	fmt.Printf("%.6f \n", pi)
	fmt.Printf("%.4f \n", pi2)

	fmt.Println("Hello World")


	base := 12 // int
	var altura int = 14 // int 
	var area int // int

	fmt.Println(base, altura, area)

	// Zero values
	var a int // 0
	var b float64 // 0
	var c string // ""
	var d bool // false

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)
}