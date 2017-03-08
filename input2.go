package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println("El texto que leyo es " + s)
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
