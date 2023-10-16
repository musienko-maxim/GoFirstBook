package main

import "fmt"

type subScriber1 struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subScriber1) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subScriber1 {
	var s subScriber1
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subScriber1) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)

}
