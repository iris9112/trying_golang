//Servidor de ficheros estaticos en una sola linea

package main

import "net/http"

func main() {
	//NOTA:  preguntar a michel porque muestra todos los archivos del directorio web y no solo la carpeta public
	http.ListenAndServe(":8088", http.FileServer(http.Dir(".")))
}
