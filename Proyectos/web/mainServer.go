package main

import (
	"fmt"
	"net/http"
)

//funcion que escribe peticion y retorna la peticion
func hola(w http.ResponseWriter, r *http.Request) {
	//acumular la respuesta y enviarla por la red
	//	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprint(w, a)
}

func main() {
	//cuando se conecte a / llama la funcion hola
	http.HandlerFunc("/", hola)

	//en q puerto vamos a escuchar
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
