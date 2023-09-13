package main

import "fmt"

//func maximum(numbers ...float64) float64 {
//	max := math.Inf(-1)
//	for _, number := range numbers {
//		if number > max {
//			max = number
//		}
//	}
//	return max
//}
//
//func main() {
//	fmt.Println(maximum(71.8, 56.2, 89.5))
//	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
//}
func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
}
