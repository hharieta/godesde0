package funciones

import "fmt"

func Exponente(n int) {
	if n > 65999 {
		return
	}
	fmt.Println(n)
	Exponente(n * 2)
}
