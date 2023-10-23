package ejercicios

import "strconv"

func Ejercicio1(s string) (int, string) {

	if num, err := strconv.Atoi(s); err != nil {
		return 1, "Error de conversión " + err.Error()
	} else if num > 100 {
		return num, "Es mayor a 100"
	} else if num < 100 {
		return num, "Es menor a 100"
	}
	return 1, "Error"
}
