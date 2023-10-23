package funciones

import (
	"fmt"
)

func Calculos() {

	calculo := func(num_1 int, num_2 int) int {
		return num_1 + num_2
	}
	fmt.Println(calculo(3, 9))
	calculo = func(num_1 int, num_2 int) int {
		return num_1 * num_2
	}
	fmt.Println(calculo(3, 9))
}
