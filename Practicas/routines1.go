package main

import (
	"fmt"
	"time"
)

var id int

func main() {
	//tienen diferentes id
	//crear un goruotines
	print(0, 10, 400)
	print(10, 20, 300)
	print(20, 30, 900)
}

func print(start int, end int, delay int) {

	id++
	for start < end {
		fmt.Printf("[%v]: %v\n", id, start)
		//Para hacer pausa entre hilos
		time.Sleep(time.Duration(delay) * time.Millisecond)
		start++
	}

}
