package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es un segundo mensaje")
}

func EjemploPanic() {
	defer func() {
        if err := recover(); err!= nil {
            log.Fatalf("Ocurrio un error que genero un panic \n %v", err)
        }
    }()
	
    VemosDefer()
	a := 1
	if a == 1 {
		panic("Se encontro el 1")
	}
}
