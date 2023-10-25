package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func NombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		// 1 segundo
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	canal1 <- true
}
