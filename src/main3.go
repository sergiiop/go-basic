package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calculate(f figuras2D) float64 {
	fmt.Println("Area", f.area())

	return f.area()
}

func main3() {
	myCuadrado := cuadrado{base: 12}

	myRectangulo := rectangulo{base: 2, altura: 3}

	areaCuadrado := calculate(myCuadrado)

	calculate(myRectangulo)

	fmt.Println(areaCuadrado)

	// Lista interface

	// [1, "bo", true]

	myInterface := []interface{}{"Hola", 12, 4.2}

	fmt.Println(myInterface...) // Se pone ... para indicar que se van a imprimir cada uno de los elementos

}