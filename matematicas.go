package main

import (
	"fmt"
	"math"
)

func main() {
	AreaCirculo(2.5)
	Multiplicacion(26, 25, 36)
	Suma(1, 2, 3, 4)
}

//AreaCirculo funcion con un parametro
func AreaCirculo(r float64) {
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf("Area del circulo (r=%.3f) = %.3f \n", r, area)
}

//Multiplicacion recibe varios parametros
func Multiplicacion(a, b, c float64) {
	result := a * b * c
	fmt.Printf("Multiplicacion de a=%.2f, b=%.2f y c=%.2f es igual a: %.2f \n", a, b, c, result)
}

/*Suma recibe cantidad de parametros indefinidos
func nombre_de_funcion tipo_de_dato_que_devuelve {
return valor_que_devuelve
}
*/
func Suma(n ...int) int {
	resultado := 0
	return resultado + n
}
