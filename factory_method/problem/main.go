package main

import "fmt"

// Ver 1:

// type Truck struct{}

// func (t *Truck) Deliver() {
// 	fmt.Println("Deliver by land in a truck")
// }

// func main() {
// 	truck := &Truck{}
// 	truck.Deliver()	// Output: Deliver by land in a truck
// }

// Ver 2:

type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Deliver by land in a truck")
}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Deliver by sea in a ship")
}

func main() {
	var transportType string = "ship"
	if transportType == "truck" {
		truck := &Truck{}
		truck.Deliver()
	} else if transportType == "ship" {
		ship := &Ship{}
		ship.Deliver()
	}
	// Output: Deliver by sea in a ship
}
