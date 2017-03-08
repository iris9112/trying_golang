//valores

package main

import "fmt"

func main() {
	//Queremos imprimir la cadena golang
	fmt.Println("go" + "lang")

	//Queremos aprender a sumar/restar enteros
	fmt.Println("5+5 =", 5+5)

	//Queremos aprender a manejar los flotantes
	fmt.Println("5/3.0 = ", 5/3.0)

	//Manejo de booleanos
	fmt.Println("resultado=")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	//Declarando la variable
	var a = "Empezar"

	//Declarando más de una variable
	var b, c int = 4, 5
	fmt.Println("a = ", a, "b = ", b, "c = ", c)

	//Declarando variables con resultado booleano
	var d = true
	fmt.Println(d)

	//otra forma de declarar la variable
	f := "short"
	var m = "cadena"
	fmt.Println(f)
	fmt.Println(m)

}
