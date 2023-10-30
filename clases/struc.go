package clase1

import "fmt"

// type para definir un tipo de dato
// struct para definir una estructura

type Car struct {
	brand string
	year  int
}

func main22() {
	myCar := Car{
		brand: "Ford",
		year:  2020,
	}

	fmt.Println(myCar)

	// Instanciar como clase vacia

	var otherCar Car

	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)
}