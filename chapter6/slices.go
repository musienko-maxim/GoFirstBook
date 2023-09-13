package main

import "fmt"

func main() {
	////	var notes []string
	////	notes = make([]string, 7)
	////	notes[0] = "do"
	////	notes[1] = "re"
	////	notes[2] = "mi"
	////	fmt.Println(notes[0])
	////	fmt.Println(notes[1])
	////	//
	////	primes := make([]int, 5)
	////	primes[0] = 2
	////	primes[1] = 3
	////	fmt.Println(primes[0])
	//notes := make([]string, 7)
	//primes := make([]int, 5)
	//fmt.Println(len(notes))
	//fmt.Println(len(primes))
	//letters := []string{"a", "b", "c"}
	//for i := 0; i < len(letters); i++ {
	//	fmt.Println(letters[i])
	//}
	//for _, letter := range letters {
	//	fmt.Println(letter)
	//}
	array1 := [5]string{"a", "b", "c,", "d", "e"}
	slice1 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice1)
}
