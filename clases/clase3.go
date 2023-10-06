package clase1

import "fmt"

func clase3() {
	// Declaración de variables

	helloMessage := "Hello"

	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)

	// printf

	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf (guardar en una variable)

	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)

	fmt.Println(message)

	// Tipo de datos

	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}