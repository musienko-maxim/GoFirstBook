package main

import "fmt"

//func main() {
//	var width, height, area float64
//	width = 4.2
//	height = 3.0
//	area = width * height
//	fmt.Printf("%.2f liters needed\n", area/10.0)
//	width = 5.2
//	height = 3.5
//	area = width * height
//	fmt.Printf("%.2f liters needed\n", area/10.0)
//}
func main() {
	repeatLine("hello", 3)
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
