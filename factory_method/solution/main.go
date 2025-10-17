package main

import "fmt"

type Transport interface {
	Deliver()
}

type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Deliver by land in a truck")
}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Deliver by sea in a ship")
}

type Logistics interface {
	CreateTransport() Transport
}

type RoadLogistics struct{}

func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

type SeaLogistics struct{}

func (s *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

func planDelivery(logistics Logistics) {
	transport := logistics.CreateTransport()
	transport.Deliver()
}

func main() {
	var logistics Logistics

	logistics = &RoadLogistics{}
	planDelivery(logistics)

	logistics = &SeaLogistics{}
	planDelivery(logistics)
}
