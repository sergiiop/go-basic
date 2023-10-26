package clase1

import "fmt"

func maps() {

	// Crear un map
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	// Recorrer un map

	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor

	// Si se accede a un valor que no esta en el diccionario se obtiene el cero value del tipo de dato de los valores, es decir como en este caso es int, se obtiene 0

	// Si encontramos un valor y le pasamos otra variable, este pasara a ser un validador si la key esta en el map
	value, ok := m["Josep"]

	fmt.Println(value, ok)
}