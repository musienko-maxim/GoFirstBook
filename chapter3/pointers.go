package main

import "fmt"

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)

}
func double(number *int) {
	*number *= 2

}

//func createPointer() *float64 {
//	var myFloat = 98.5
//	return &myFloat
//}
//
//func main() {
//	var myFloatPointer *float64 = createPointer()
//	fmt.Println(*myFloatPointer)
//
//}
