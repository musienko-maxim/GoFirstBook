package main

import "fmt"

type Coordinates struct {
	Latitude  float64
	Longitude float64
}
type Lamdmark struct {
	Coordinates
	Name string
}

func main() {
	location := Lamdmark{}
	location.Name = "The Googleplex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
