package main

import "fmt"

//El usuario representa un usuario del sistema
type usuario struct {
	nombre string
	email  string
}

//NuevoUsuario crea y regresa apuntadores de valores de tipo usuario
func nuevoUsuario() (*usuario, error) {
	return &usuario{"Isabel", "iris9112@gmail.com"}, nil
}

func main() {

	//Crear valores del tipo usuario
	u, err := nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Imprimir el valor U
	fmt.Println(*u)

	//2do llamado a la función y que solo verifique el error en el regreso

	_, err = nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}
}
