package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgumentFunction(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, b int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello World!")
	tripleArgumentFunction(1, 2, "3")

	value := returnValue(2)

	fmt.Println(value)

	value1, _ := doubleReturn(2)

	fmt.Println(value1)
}

func calculateArea(width, height int) int {
	return width * height
}

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}