package main

import (
	e "github.com/agudelozc/godesde0/ejer_interfaces"
	models "github.com/agudelozc/godesde0/modelos"
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
	}

	num, text := ejercicios.Ejercicio1("99")
	fmt.Println("El numero ingresado es =", num, "Y", text)

	teclado.IngresoNumeros()

	iteraciones.Iterar()
	fmt.Println(ejercicios.Multiplicar())

	files.GrabaTabla()

	files.SumaTabla()
	files.LeoArchivo()
	funciones.Calculos()
	funciones.LlamarClosure()
	funciones.Exponencia(2)
	arreglos_slices.Capacidad()
	mapas.MostrarMapas()
	users.AltaUsuario()*/
	Pedro := new(models.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(models.Mujer)
	e.HumanosVivo(Maria)
	//e.HumanosRespirando(Maria)
}
