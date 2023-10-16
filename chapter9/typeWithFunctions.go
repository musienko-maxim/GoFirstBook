package main

import "fmt"

//type MyType string
//
//func (m MyType) sayHi() {
//	fmt.Println("Hi from", m)
//
//}
//
//func main() {
//	value := MyType("a MyType value")
//	value.sayHi()
//	anotherValue := MyType("another value")
//	anotherValue.sayHi()
//}
type Number1 int

func (n Number1) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number1) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	ten := Number1(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number1(4)
	four.Add(3)
	four.Subtract(2)
}
