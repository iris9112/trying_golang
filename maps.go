//Mapas

package main

import "fmt"

func main() {

	//declaracion del mapa
	departamentos := make(map[string]int)

	//inicilizar los datos en el mapa
	departamentos["Devs"] = 25
	departamentos["Marketing"] = 50
	departamentos["Ejecutivos"] = 4
	departamentos["ventas"] = 60
	departamentos["Mantenimiento"] = 8

	//Desplegar con iteraciones el valor de cada par llave/valor

	for key, value := range departamentos {

		fmt.Printf("Deps: %s Personas %d\n", key, value)
	}

	//recorriendo tipos de datos compuestos con range
	numeros := []int{2, 4, 6}
	suma := 0
	for _, numero := range numeros {
		suma += numero
	}

	fmt.Println("suma: ", suma)

	for i, numero := range numeros {
		if numero == 6 {
			fmt.Println("index: ", i)
		}

	}

	//mapa
	//mapa := map[string]string{"a": auto, "b": bebe, "c": casa}

}
