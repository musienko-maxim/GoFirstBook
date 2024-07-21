package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck_ string

func (t Truck_) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck_) Brake() {
	fmt.Println("Stopping")
}
func (t Truck_) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck_) LoadCargo(cargo string) {
}

type Vehicle2 interface {
	Accelerate()
	Steer(direction string)
	Brake()
}

func main() {
	var vehicle Vehicle2 = Car("Toyoda Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck_("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
}
