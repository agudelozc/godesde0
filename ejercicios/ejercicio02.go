package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese el numero al que desea la tabla del 1 al 10: ")
		if scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue

			} else {
				for i := 0; i <= 10; i++ {
					fmt.Println(num * i)
				}
			}
		}
	}
}
