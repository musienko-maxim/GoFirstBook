package main

import (
	"fmt"
)

var metersPerLiter float64

//func paintNeeded(width float64, height float64) float64 {
//	area := width * height
//	return area / metersPerLiter
//}

func main() {
	err := fmt.Errorf("a height of %0.2f is invalid", -2.333333)
	fmt.Println(err.Error())
	fmt.Println(err)
	//log.Fatal(err)
	//metersPerLiter = 10.0
	//amount := paintNeeded(4.2, -3.0)
	//fmt.Printf("%0.2f liters needed\n", amount)
}
