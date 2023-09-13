package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func main() {

	status("Alma")
	status("Carl")
	status("Rohit")
	//counters := map[string]int{"a": 3, "b": 0}
	//var value int
	//var ok bool
	//value, ok = counters["a"]
	//fmt.Println(value, ok)
	//value, ok = counters["b"]
	//fmt.Println(value, ok)
	//value, ok = counters["c"]
	//fmt.Println(value, ok)
	//fmt.Println(ranks["bronze"])
	//fmt.Println(ranks["gold"])
	//elements := make(map[string]string)
	//elements["H"] = "Hydrogen"
	//elements["Li"] = "Lithium"
	//fmt.Println(elements["Li"])
	//fmt.Println(elements["H"])
	//myMap := map[string]float64{"a": 1.2, "b": 5.6}
	//ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	//fmt.Println(ranks["gold"])
	//fmt.Println(ranks["bronze"])

	//elements := map[string]string{
	//	"H":  "Hydrogen",
	//	"Li": "Lithium",
	//}
	//fmt.Println(elements["H"])
	//fmt.Println(elements["Li"])
	//
	//numbers := make(map[string]int)
	//numbers["I've been assigned"] = 12
	//fmt.Printf("%#v\n", numbers["I've been assigned"])
	//fmt.Printf("%#v\n", numbers["I haven't been assigned"])
	//var myMap map[int]string = make(map[int]string)
	//myMap[3] = "three"
	//fmt.Printf("%#v\n", myMap)
}
