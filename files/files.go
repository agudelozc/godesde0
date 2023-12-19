package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/agudelozc/godesde0/ejercicios"
)

var fileName string = "./files/txt/Tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Multiplicar()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file" + err.Error())
		return
	}
	fmt.Fprintln(file, texto)
	file.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Multiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Error concatenando contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
