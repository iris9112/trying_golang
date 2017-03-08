// Interfaces

// Declarar una interface llamada speaker con un método llamado speak.
// Declarar un struct llamado ingles que represente a una persona que hable Inglés.
// Declarar un struct llamado chino que represente a una persona que hable chino.
// Implementar la interface speaker para cada struct usando un valor y strings: "Hello World" y "你好世界"
// Declarar una variable del tipo speaker y asignarle una dirección de tipo Inglés y llamar el método.
// Hay que hacerlo para chino.

package main

import "fmt"

// speaker implementa lo que dice una persona
type speaker interface {
	speak()
}

type listener interface {
	listen()
}

//interfaz q es composicion de las anteriores
type speakerListener interface {
	listener
	speaker
}

// inglés representa a una persona que habla inglés

type ingles struct {
	name string
}

// speak implementa la interface speaker
func (ingles) speak() {
	fmt.Println("Hello World")
}

func (i ingles) listen() {
	fmt.Printf("%s listen in english \n", i.name)
}

// chino representa a una persona que habla chino

type chino struct {
	name string
}

// speak implementa la interface speaker
func (chino) speak() {
	fmt.Println("你好世界")
}
func (chino) listen() {
	fmt.Println("你你你你你你你你你你好世界")
}

func main() {
	//Declarar una variable para interface
	var sp speaker
	var lp listener
	var sl speakerListener

	// Asignarle un valor a la interface y vamos a llamar al método (de la interface)
	var i = ingles{name: "Isa"}
	sp = i
	sp.speak()
	lp = i
	lp.listen()
	fmt.Println("El ingles es: ", i.name)

	// Chino
	var c = chino{name: "ho chimin"}
	sl = c
	sl.speak()
	sl.listen()

	// Crear nuevos valores y llamar la función
	decirHola(new(ingles))
	decirHola(&chino{})
}

// decirHola abstrae la función de hablar

func decirHola(sp speaker) {
	sp.speak()
}
