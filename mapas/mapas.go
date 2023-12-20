package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Colombia"] = "Bogota"
	paises["Peru"] = "Lima"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Liga 1": 1,
		"Liga 2": 2,
		"Liga 3": 3,
	}
	fmt.Println(campeonato)

	for _, p := range campeonato {
		fmt.Println(p)
	}
	delete(campeonato, "Liga 1")
	fmt.Println(campeonato)

	p, e := campeonato["Liga 2"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", p, e)
}
