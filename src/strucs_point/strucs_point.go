package strucspoint

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50

	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	// & -> Devuelve la dirección de memoria de una variable
	// * -> Devuelve el valor de una variable a partir de su dirección de memoria

	myPc := pc{ram: 16, disk: 200, brand: "Asus"}

	fmt.Println(myPc)
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}