package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir := "/home/isabel/golibs/src/github.com/iris9112/readfiles/textos"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(len(files))

	var str string

	for i := 0; i < len(files); i++ {
		file := files[i].Name()
		fmt.Println(file)

		fileread, _ := ioutil.ReadFile(dir + string(os.PathSeparator) + file)

		str = str + string(fileread)
		fmt.Println(fileread)
	}

	fmt.Println(str)
	err = ioutil.WriteFile("salida.txt", []byte(str), 0644)
	if err != nil {
		panic(err)
	}

}
