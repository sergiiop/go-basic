package clase1

import "fmt"

func clase2() {
	// Area cuadrado
	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)

	x:= 10
	y:= 50

	// suma
	result := x + y

	fmt.Println("Suma:", result)

	// resta
	result = y - x // como ya esta definida la variable result, no se usa :=

	fmt.Println("Resta:", result)

	// multiplicación

	result = x * y

	fmt.Println("Multiplicación:", result)

	// division

	result = y / x

	fmt.Println("Division:", result)

	// modulo

	result = y % x

	fmt.Println("Modulo:", result)

	// incremental
	x++
	fmt.Println("Incremental:", x)

	// decremental
	x--
	fmt.Println("Decremental:", x)

	// Reto
	// rectángulo, trapecio y circulo

	// rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("Area rectángulo:", areaRectangulo)

	// trapecio
	baseMayorTrapecio := 20
	baseMenorTrapecio := 10
	alturaTrapecio := 5

	areaTrapecio := ((baseMayorTrapecio + baseMenorTrapecio) / 2) * alturaTrapecio

	fmt.Println("Area trapecio:", areaTrapecio)

	// circulo
	const pi float64 = 3.14159265
	radioCirculo := 10

	areaCirculo := pi * float64(radioCirculo * radioCirculo)

	fmt.Println("Area circulo:", areaCirculo)
}