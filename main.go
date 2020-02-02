package main

import (
	"fmt"
	"io/ioutil"
)

func main () {
	var insBytes, err = ioutil.ReadFile("./program.bf")
	if err != nil {
		panic(err)
	}

	var vm = newVm(string(insBytes))
	vm.interpretAll()

	fmt.Println("Program fin.")
}
