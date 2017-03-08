package main

import "fmt"

func print(chann chan int) {
	var x int
	//creamos un for infinito??
	for {
		x = <-chann
		fmt.Printf("%v\n", x)
	}

}

func main() {
	//crear un canal
	channel := make(chan int)
	go print(channel)

	i := 0
	for {
		i++
		channel <- 1
	}

}
