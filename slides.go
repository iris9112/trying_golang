// Slices

package main

import "fmt"

func main() {

	// Declarar un slice de enteros vacío.
	var numeros []int

	// Crear un loop que meta 10 valores al slice.
	for i := 0; i < 10; i++ {
		numeros = append(numeros, i*10)
	}
	// Iterar sobre el slice y mostrar cada valor.
	for _, numero := range numeros {
		fmt.Println(numero)
	}

	//Declarar un slice de 5 strings e inicializar dicho slice con valores string.
	frutas := []string{"manzana", "naranja", "pera", "mango", "coco"}

	//mostrar cada indice y nombre de las frutas
	for i, fruta := range frutas {
		fmt.Printf("Index: %d Fruta: %s\n", i, fruta)
	}

	// Tomar un slice del primer y segundo índice

	slice := frutas[1:3]

	// Desplegar la posición y el valor de cada uno de estos elementos en el nuevo slice.
	for i, fruta := range slice {
		fmt.Printf("Index2: %d Fruta2: %s\n", i, fruta)
	}
}
