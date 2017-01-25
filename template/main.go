package main

import (
	"os"
	"text/template"
)

//Persona estructura con nombre y edad
type Persona struct {
	Nombre string
	Edad   int
}

//para procesar el template
const tp = `Nombre: {{.Nombre}} Edad: {{.Edad}}`

func main() {
	isa := Persona{"ISabel", 25}
	t := template.New("isa")
	t, err := t.Parse(tp)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, isa)
	if err != nil {
		panic(err)
	}
}
