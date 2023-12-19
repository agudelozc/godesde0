package main

import (
	"fmt"

	"github.com/agudelozc/godesde0/ejercicios"
)

func main() {
	//variables.MuestroEnterosa()
	//estado, texto := variables.ConviertoaTexto(1588)
	//fmt.Println(estado, texto)

	//os := runtime.GOOS

	/*if os := runtime.GOOS; os == "windows" {
		fmt.Println("Estamos en windows")
	} else {
		fmt.Println("Estamos en otro sistema")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Estamos en windows")
	case "linux":
		fmt.Println("Estamos en linux")
	default:
		fmt.Printf("%s \n", os)
	} */

	num, text := ejercicios.Ejercicio1("99")
	fmt.Println("El numero ingresado es =", num, "Y", text)
}
