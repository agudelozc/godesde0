package ejer_interfaces

import (
	"fmt"

	"github.com/agudelozc/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}

func HumanosVivo(hu interfaces.Humano) {
	hu.EstaVivo()
	fmt.Printf("Soy un/a %s, y estoy vivo \n", hu.Sexo())
}
