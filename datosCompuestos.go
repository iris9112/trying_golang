package main

import "fmt"

//Declarar un array de 5 string cada uno, los vamos a inicializar con 0
//Declarar un segundo arreglo de 5 strings, los vamos a inicializar con los valores de los strings.
//Asignar el segundo arreglo al primero y desplegar los resultados del primero
//Mostrar el valor de la cadena y la dirección de cada elemento

func main() {
	//Declarar arreglos de String que guarden los nombres
	var nombres [5]string

	//Declarar un arreglo pre-llenado con nombres de amigos
	amigos := [5]string{"kathe", "karo", "angel", "root", "leo"}

	//Asignar el arreglo amigos a nombres
	nombres = amigos

	//imprimir nombres con for
	for i, nombre := range nombres {
		fmt.Println(nombre, &nombres[i])
	}

	alumnos := 1
	for alumnos <= 3 {
		fmt.Println("alumno no: ")
		fmt.Println(alumnos)
		alumnos = alumnos + 1
	}

	for calificaciones := 7; calificaciones <= 9; calificaciones++ {
		fmt.Println(calificaciones)
	}

}
