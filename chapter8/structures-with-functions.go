package main

import "fmt"

type subbcriber1 struct {
	name   string
	rate   float64
	active bool
}

func printINfo(s subbcriber1) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubScriber(name string) subbcriber1 {
	var s subbcriber1
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func main() {
	subscriber1 := defaultSubScriber("Aman Singh")
	subscriber1.rate = 4.99
	printINfo(subscriber1)

	subscriber2 := defaultSubScriber("Beth Ryan")
	printINfo(subscriber2)
}
