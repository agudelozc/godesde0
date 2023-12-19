package main

import (
	"fmt"

	"github.com/agudelozc/godesde0/variables"
)

func main() {
	//variables.MuestroEnterosa()
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado, texto)
}
