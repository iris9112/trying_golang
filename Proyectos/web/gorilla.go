package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//GETUsers crea recursos
func GETUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje del metodo Get")
}

//POSTUsers envia recursos
func POSTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje del metodo post")
}

//PUTUsers actualiza recursos
func PUTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje del metodo put")
}

//DELETEUsers elimina recursos
func DELETEUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje del metodo delete")
}

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GETUsers).Methods("GET")
	r.HandleFunc("/api/users", POSTUsers).Methods("POST")
	r.HandleFunc("/api/users", PUTUsers).Methods("PUT")
	r.HandleFunc("/api/users", DELETEUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Conectandose a //localhost:8080 ...")
	log.Fatal(server.ListenAndServe())

}

//Para testear APIs en Google chrome utilizar Postman
