package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/user/{name:[a-z]+}/profile", profile).Methods("GET")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	//muestra en la pagina la cadena
	w.Write([]byte("Hello user " + name))
}

//localhost:3000/user/isa/profile
