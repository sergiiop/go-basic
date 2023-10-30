package main

import "fmt"

func main2() {
	// array

	var array [4]int

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4

	fmt.Println(array, len(array), cap(array))

	// slices

	var slice = []int  {
		1,
		2,
		3,
		4,
		5,
		6,
		7,
	}

	defer fmt.Println(slice)

	fmt.Println(slice, len(slice), cap(slice))

	// Métodos de slices

	fmt.Println(slice[0])

	fmt.Println(slice[:3])

	fmt.Println(slice[2:])

	fmt.Println(slice[2:4])

	// append

	slice = append(slice, 8)

	fmt.Println(slice, len(slice), cap(slice))
	newSlice := []int{9, 10, 11}


	slice = append(slice, newSlice...)

	fmt.Println(slice)




}