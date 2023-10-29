package main

import (
	// "github.com/hharieta/godesde0/arreglos"
	// e "github.com/hharieta/godesde0/ejerciciosinter"
	// "github.com/hharieta/godesde0/modelos"
	//"github.com/hharieta/godesde0/deferpanic"
	// "fmt"

	// "github.com/hharieta/godesde0/goroutines"
	"github.com/hharieta/godesde0/webserver"
)

func main() {
	// fmt.Println("Hello, World!")
	// variables.MostrarEnteros()
	// variables.OtrasVariables()
	// estado, texto := variables.Convert2Text(4)

	// fmt.Println("Estado: ", estado, "Texto: ", texto)

	// if os := runtime.GOOS; os == "Linux" || os == "darwin" {
	// 	fmt.Println("Esto no es Windows: ", os, "god")
	// } else {
	// 	fmt.Println("Esto es un Windowszzz")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// n, s := ejercicios.Ejercicio1("uih")

	// fmt.Println(n, s)

	// funciones.Calculos()
	// funciones.LlamarClosure()
	// funciones.Exponente(2)

	// arreglos.MostrarArray()
	// arreglos.IngresarDatosArray(8, 3)
	// arreglos.MostrarArray()
	// arreglos.IterarArray()

	// arreglos.Slices()
	// arreglos.CapacidadSlices()
	// arreglos.RellenarSlice()

	// maps.MostrarMapas()

	//usuarios.AltaUsuario()

	// Gatovsky := new(modelos.Hombre)
	// Mary := new(modelos.Mujer)

	// e.HumanoRespirando(Gatovsky)
	// e.HumanoRespirando(Mary)

	// deferpanic.VerDefer()
	// deferpanic.EjemploPanic()

	// canal1 := make(chan bool)
	// go goroutines.NombreLento("Gatovsky", canal1)

	// defer func() {
	// 	<-canal1
	// 	fmt.Println("Fin gourutine!!")
	// }()

	webserver.Webserver()
}
