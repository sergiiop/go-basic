package clase1

import (
	"fmt"
	"log"
	"strconv"
)

func forIf() {

	// For condicional

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 10; i > 1; i-- {
		fmt.Println(i)
	}

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es 1 y 2")
	}

	// convertir de string a int

	value, err := strconv.Atoi("45")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}