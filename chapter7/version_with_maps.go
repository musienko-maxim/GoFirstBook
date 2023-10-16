package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {
	//grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	//for name, _ := range grades {
	//	fmt.Println("***********", name)
	//}
	lines, err := datafile.GetStrings("/home/mmusiien/tmp/voutes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

}
