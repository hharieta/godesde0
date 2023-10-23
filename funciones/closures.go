package funciones

import "fmt"

func tabla(valor int) func() int {
	n := valor
	secucencia := 0

	return func() int {
		secucencia++
		return n * secucencia
	}
}

func LlamarClosure() {
	tabla_num := 2
	MiTabla := tabla(tabla_num)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
