package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//Para leer archivos se necesita verificar los llamados de error
func check(e error) {
	if e != nil {
		return
	}
}

func main() {

	ruta := "/home/isabel/curso_go/Proyectos/reading/file"

	//poner el contenido en memoria
	dat, err := ioutil.ReadFile(ruta)
	check(err)
	fmt.Print(string(dat))

	//abrir fichero
	f, err := os.Open(ruta)
	check(err)

	//leer bytes
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("Forma1: %d bytes: %s\n", n1, string(b1))

	//tambien puedes conocer la ubicacion en el archivo y leer desde ahí
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("Forma2: %d bytes @ %d: %s\n", n2, o2, string(b2))

	//Leer con ReadAtLeast es más robusto
	o3, err := f.Seek(7, 0)
	check(err)
	b3 := make([]byte, 4)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("Forma3: %d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(500)
	check(err)
	fmt.Printf("Forma 4: n bytes: %s\n", string(b4))

	f.Close()

}

/*
Hello = contenido del archiov
dat = nombre del archivo
/tmp = ruta
$ echo "hello" > /tmp/dat

llamar al archivo
$ go run readingFiles.go

crear archivo en directorio actual:
(Con pwd se ve la ruta completa)
echo "abcdefg hijklmn opqrs" > /home/isabel/curso_go/Proyectos/reading/file
*/
