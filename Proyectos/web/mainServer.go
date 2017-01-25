package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func handlerPrueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bienvenido a la sesión de Pruebas</h1>")
}

func handlerUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bienvenido Usuario</h1>")
}

func main() {

	//creamos el manejador como estructura y no como funcion
	msg := mensaje{
		msg: "<h1>Bienvenido a la seccion de manejadores</h1>",
	}

	//Servir archivos estaticos, recibe el directorio donde se ubican los archivos
	//Para eso se llama a la funcion dir q toma el nombre del directorio o la ruta
	fs := http.FileServer(http.Dir("public"))

	//ServerMux es el enrutador del paquete http
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("/prueba", handlerPrueba)
	mux.HandleFunc("/usuario", handlerUsuario)
	//como ya tenemos una struct manejador entonces no necesitamos pasarla
	//como funcion con HandlerFunc sino directamente llamamos al manejador
	mux.Handle("/manejador", msg)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Conectandose a //localhost:8080 ...")
	log.Fatal(server.ListenAndServe())

}
