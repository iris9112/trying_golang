package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("Hello, I am ISa")
	}
}

//Para entrar al primer bloque, ejecutar:
//go run condictionals.go "Into the tunnel"
