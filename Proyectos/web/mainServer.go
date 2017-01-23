package main

import (
	"fmt"
	"net/http"
)

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
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

func main() {

	//creamos el manejador como estructura y no como funcion
	msg := mensaje{
		msg: "Hola mundo de nuevo",
	}

	//ServerMux es el enrutador del paquete http
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/prueba", handlerPrueba)
	mux.HandleFunc("/usuario", handlerUsuario)

	//como ya tenemos una struct manejador entonces no necesitamos pasarla
	//como funcion con HandlerFunc sino directamente llamamos al manejador
	mux.Handle("/manejador", msg)

	//Se le pasa el puerto y un enrutador
	http.ListenAndServe(":8080", mux)

}
