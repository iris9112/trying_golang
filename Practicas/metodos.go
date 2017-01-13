//metodos

//struct camionero: nombre, carro, trasteos

//metodo para calcular total de trasteos

//slice con varios camioneros
//iterar sobre camioneros y total de trasteos

package main

import "fmt"

//camionero
type camionero struct {
	nombre   string
	carro    string
	trasteos int
	dias     int
}

//total trasteos
func (c *camionero) totalTrasteos() int {
	return c.trasteos * c.dias
}

func main() {

	camioneros := []camionero{
		{"Felipe", "van", 49, 20},
		{"Efra", "furgon", 60, 30},
		{"Jorge", "camioneta", 34, 15},
		{"Angel", "camioneta", 60, 30},
		{"tiven", "van", 15, 3},
	}

	//mostrar total de trasteos por camionero
	for _, camionero := range camioneros {
		fmt.Println(camionero.nombre, camionero.totalTrasteos())
	}

}
