package main

import (
	"fmt"
	"net/http"
)

func main() {

	//ServerMux es el enrutador del paquete http
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/prueba", handlerPrueba)
	mux.HandleFunc("/usuario", handlerUsuario)

	//Se le pasa el puerto y un enrutador
	http.ListenAndServe(":8080", mux)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo</h1>")
}

func handlerPrueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola mundo desde /Pruebas</h1>")
}

func handlerUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bienvenido Usuario</h1>")
}
