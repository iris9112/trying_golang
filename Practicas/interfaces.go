//interfaces

package main

import "fmt"

//Speaker implementa lo que dice una persona
type speaker interface {
	speak()
}

//Ingles representa una persona que habla Ingles
type ingles struct{}

//speak implementa la interface Speaker
func (ingles) speak() {
	fmt.Println("Hello world")
}

type chino struct{}

func (chino) speak() {
	fmt.Println("Hello world")
}
