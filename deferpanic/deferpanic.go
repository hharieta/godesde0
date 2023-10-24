package deferpanic

import (
	"fmt"
	"log"
)

func VerDefer() {
	fmt.Println("Primer Mensaje")

	defer fmt.Println("Mensaje Final")

	fmt.Println("Segundo Mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Panic Erro at: %v \n", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("PANIC!!!")
	}
}
