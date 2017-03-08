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
//se agrega el range para recorrer el slide de estructura
const tp = `
  {{range .}}
    {{if .Edad}}
      Nombre: {{.Nombre}} Edad: {{.Edad}} - Correcto
      {{else}}
      Nombre: {{.Nombre}} Edad: {{.Edad}} - INCorrecto
    {{end}}
  {{end}}`

func main() {
	grupo := []Persona{
		{"Johaneto", 30},
		{"Erika", 22},
		{"Michel", 29},
		{"Angel", 26},
		{"Lucy", 0}, //no cumple la condicion del if
	}
	t := template.New("grupo")
	t, err := t.Parse(tp)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, grupo)
	if err != nil {
		panic(err)
	}
}
