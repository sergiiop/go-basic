package struct_modify

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

// type para definir un tipo de dato
// struct para definir una estructura

type Car struct {
	brand string
	year  int
}

func main2() {
	var myCar pk.CarPublic

	myCar.Brand = "Ferrari"

	fmt.Println(myCar)

	pk.PrintMessage("Hola Platzi")
}