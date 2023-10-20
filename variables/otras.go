package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Fecha time.Time
var Estado bool
var Nombre string
var Sueldo float32

func OtrasVariables() {

	Fecha = time.Now()
	Estado = true
	Nombre = "Gatovsky"
	Sueldo = 5.00

	fmt.Println("Fecha: ", Fecha)
	fmt.Println("Estado: ", Estado)
	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Sueldo: ", Sueldo)

}

func Convert2Text(n int) (bool, string) {

	texto := strconv.Itoa(n)

	return true, texto
}
