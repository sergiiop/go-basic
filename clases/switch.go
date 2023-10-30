package clase1

import "fmt"

func switch2() {
	switch moduol := 5 % 2; moduol {
	case 0:
		fmt.Println("Es par")

	default:
		fmt.Println("Es impar")
	}

	value := 100

	switch {
		case value > 100:
			fmt.Println("Es mayor a 100")
		case value < 0:
			fmt.Println("Es menor a 0")
		default:
			fmt.Println("No condición")
	}
}