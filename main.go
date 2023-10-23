package main

import (
	"fmt"
	"runtime"
)

func main() {
	// fmt.Println("Hello, World!")
	// variables.MostrarEnteros()
	// variables.OtrasVariables()
	// estado, texto := variables.Convert2Text(4)

	// fmt.Println("Estado: ", estado, "Texto: ", texto)

	if os := runtime.GOOS; os == "Linux" || os == "darwin" {
		fmt.Println("Esto no es Windows: ", os, "god")
	} else {
		fmt.Println("Esto es un Windowszzz")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
