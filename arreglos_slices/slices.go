package arreglos_slices

import "fmt"

var tabla3 []int = []int{1, 2, 3}
var arreglo [10]int = [10]int{1, 4, 55, 66, 77, 88}

func MuestroSlice() {
	fmt.Println(tabla3)
	porcion := arreglo[3:]   //slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  //slice creado desde un vector, desde la posicion 0 hasta 5
	porcion3 := arreglo[6:7] //slice creado desde la posicion 6 hasta 7
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elements := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
