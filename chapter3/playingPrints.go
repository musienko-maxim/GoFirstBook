package main

import "fmt"

func main() {
	//fmt.Println("About one-third:", 1.0/3.0)
	// fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
	//fmt.Printf("A float: %f\n", 3.1415)
	//fmt.Printf("An integer: %d\n", 15)
	//fmt.Printf("A string: %s\n", "hello")
	//fmt.Printf("A boolean: %t\n", false)
	//fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	//fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	//fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	//fmt.Printf("Percent sign: %%\n")
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
	//fmt.Printf("%.2f liters needed\n", 1.8199999999999998)
}
