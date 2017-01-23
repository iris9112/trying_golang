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

	//poner el contenido en memoria
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	//abrir fichero
	f, err := os.Open("/tmp/dat")
	check(err)
	defer f.Close() //ejecuta la funcion al final

	//leer bytes
	b1 := make([]byte, 15)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//tambien puedes conocer la ubicacion en el archivo y leer desde ahí
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	//Leer con ReadAtLeast es más robusto
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(15)
	check(err)
	fmt.Printf("15 bytes: %s\n", string(b4))

}

/*

$ echo "hello hola ola" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run reading-files.go

*/
