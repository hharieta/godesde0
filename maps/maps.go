package maps

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	fmt.Println(paises)

	paises["Argentina"] = "Buenos Aires"
	paises["México"] = "CDMX"
	paises["España"] = "Madrid"

	fmt.Println(paises)
	fmt.Println(paises["México"])

	campeonato := map[string]int{
		"Argentina": 28,
		"México":    32,
		"España":    32,
	}

	fmt.Println(campeonato)
	for pais, puntaje := range campeonato {
		fmt.Println(pais, puntaje)
	}

	delete(campeonato, "España")
	fmt.Println(campeonato)

	// un map deveulve un valor y un booleano
	puntaje, existe := campeonato["España"]

	fmt.Printf("Puntaje: %d País: %t \n", puntaje, existe)
}
