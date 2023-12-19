package ejercicios

import (
	"strconv"
)

func Ejercicio1(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error en la conversión" + err.Error()
	}

	var text string
	if num > 100 {
		text = "Es mayor a 100"
	} else {
		text = "Es menor a 100"
	}
	return num, text
}
