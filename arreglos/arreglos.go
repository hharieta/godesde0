package arreglos

import "fmt"

var tabla [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MostrarArray() {
	fmt.Println(tabla)
}

func IngresarDatosArray(n int, i int) {
	tabla[i] = n
}

func IterarArray() {
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
}
