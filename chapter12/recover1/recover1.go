package main

import "fmt"

func calmDown() {
	recover()
}

func freekOut() {
	defer calmDown()
	panic("oh no")
}
func main() {
	freekOut()
	fmt.Println("Exiting normally")
}
