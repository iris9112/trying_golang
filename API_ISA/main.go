//Este API no incluye paginacion ni seguridad
//Devuelve en json todas las notas almacenadas en nuestro map
//Probar las peticiones con Postman
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Note estructura y paseo a json para hacerlo mas legible
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}

//utilizamos este mapa como BD y el indice es un string
var noteStore = make(map[string]Note)

var id int

//GetNoteHandler GET - /api/note
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	//crear un slide que va contener la nota que luego sera codificada a json
	var notes []Note

	for _, value := range noteStore {
		notes = append(notes, value)
	}
	//cabecera http para indicarle al usuario que devuelve un contenido json
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		//esto no se hace en producción
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	//pasar el cuerpo q es un json
	w.Write(j)
}

//PostNoteHandler manejador cuando vamos a crear una nueva nota - /api/note
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	//objeto nota que se debe rellenar con el json
	var note Note
	//Json q va decodificar en el body:Cuerpo de la peticion del usuario
	// se asigna a la variable err porque podria generar un error
	//si por ej el usuario envia una estructura json invalida
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	//agregar fecha de creacion
	note.CreatedAt = time.Now()
	//incrementar el id
	id++
	//crear un indice valido para nuestro mapa
	k := strconv.Itoa(id)
	noteStore[k] = note

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

//PutNoteHandler para actualizar una nota PUT - /api/note
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	//extraer el id de la nota q se va actualizar
	vars := mux.Vars(r)
	//se pasa el id como string porque así es como esta en el map
	k := vars["id"]
	//obtener los datos q el usuario nos manda
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		panic(err)
	}
	//confirmar que hace ese ok
	if note, ok := noteStore[k]; ok {
		noteUpdate.CreatedAt = note.CreatedAt
		delete(noteStore, k)
		noteStore[k] = noteUpdate
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

//DeleteNoteHandler DELETE - /api/note
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)

}

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	//estructura server
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Conectandose a http://localhost:8080 ...")
	server.ListenAndServe()
	fmt.Print(id)

}
