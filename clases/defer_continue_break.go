package clase1

import (
	"fmt"
)

func defer_continue_break() {
	// defer

	defer fmt.Println("Hola")

	fmt.Println("Mundo")

	// explicación defer
	// 1. Se ejecuta la función
	// 2. Se ejecuta el defer
	// 3. Se ejecuta el return
	// 4. Se ejecuta el defer

	// continue y break

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue

			// Cuando i es igual a 2, se salta el continue y sigue con la siguiente iteración
		}

		if i == 8 {
			break
			// Cuando i es igual a 8, se sale del for
		}

		fmt.Println(i)
	}
}