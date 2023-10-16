package main

import "fmt"

func main() {
	type subscriber struct {
		rate   float64
		name   string
		active bool
	}
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.name)
	//
	//var pet struct {
	//	name string
	//	age  int
	//}
	//pet.name = "Max"
	//pet.age = 5
	//fmt.Println("Name:", pet.name)
	//fmt.Println("Age:", pet.age)
	//subScriber1 := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	//fmt.Println("Name:")
	//fmt.Println("Name:", subScriber1.Name)
	//
	//fmt.Println("Rate:", subScriber1.Rate)
	//fmt.Println("Active:", subScriber1.Active)}
}
