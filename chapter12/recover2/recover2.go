package main

import "fmt"

func calmDawn() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDawn()
	err := fmt.Errorf("there's an error")
	panic(err)
}
