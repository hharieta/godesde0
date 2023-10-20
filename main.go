package main

import (
	"fmt"

	"github.com/hharieta/godesde0/variables"
)

func main() {
	fmt.Println("Hello, World!")
	variables.MostrarEnteros()
	variables.OtrasVariables()
	estado, texto := variables.Convert2Text(4)

	fmt.Println("Estado: ", estado, "Texto: ", texto)
}
