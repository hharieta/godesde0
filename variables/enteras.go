package variables

import "fmt"

func MostrarEnteros() {
	var entero int
	entero32 := int32(10)
	entero64 := int64(100)

	fmt.Println("Entero común = ", entero)
	fmt.Println("Entero 32 bits = ", entero32)
	fmt.Println("Entero 64 bits = ", entero64)
}
