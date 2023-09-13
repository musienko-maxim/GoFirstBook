package main

import "fmt"

var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

var primes [5]int = [5]int{2, 3, 5, 7, 11}

//func main() {
//	//fmt.Printf("%#v\n", notes)
//	//fmt.Printf("%#v\n", primes)
//	//fmt.Println(len(notes))
//	//fmt.Println(len(primes))
//	for index, note := range notes {
//		fmt.Println(index, note)
//	}
//}

func main() {
	//for index, _ := range notes {
	//	fmt.Println(index)
	//}
	for _, note := range notes {
		fmt.Println(note)
	}
}
