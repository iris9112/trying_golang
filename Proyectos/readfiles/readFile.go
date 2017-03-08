package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func concatFiles(dir string, output string) error {

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
	err = ioutil.WriteFile(output, []byte(str), 0644)
	if err != nil {
		panic(err)
	}

	return nil

}
