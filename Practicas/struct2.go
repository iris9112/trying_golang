//Structs

package main

import "fmt"

type usuario struct {
	nombre    string
	direccion string
	edad      int
}

func main() {
	//Declarar la variable usuario

	isa := usuario{
		nombre:    "Isabel",
		direccion: "calle 61 # 3A-15",
		edad:      25,
	}

	mrkaspa := usuario{
		nombre:    "Michel",
		direccion: "calle 116 #18-15",
		edad:      27,
	}

	//imprimir datos
	fmt.Println("Nombres: ", isa.nombre, mrkaspa.nombre)
	fmt.Println("Direcciones: ", isa.direccion, mrkaspa.direccion)
	fmt.Println("Edades: ", isa.edad, mrkaspa.edad)

	//otra forma de declarar Structs
	angel := struct {
		nombre    string
		direccion string
		edad      int
		hijos     int
	}{
		nombre:    "Angel",
		direccion: "",
		edad:      26,
		hijos:     1,
	}

	angel.direccion = isa.direccion

	fmt.Println("Angel:", angel.nombre, angel.direccion, angel.edad, angel.hijos)

}
