package main

import "fmt"

type House struct {
	Walls            int
	Floor            bool
	Door             bool
	Windows          int
	Roof             bool
	Backyard         bool
	SwimmingPool     bool
	HeatingSystem    bool
	Plumbing         bool
	ElectricalWiring bool
	PorchStyle       string
}

func NewHouse(
	walls int,
	floor bool,
	door bool,
	windows int,
	roof bool,
	backyard bool,
	swimmingPool bool,
	heatingSystem bool,
	plumbing bool,
	electricalWiring bool,
	porchStyle string,
) House {
	return House{
		Walls:            walls,
		Floor:            floor,
		Door:             door,
		Windows:          windows,
		Roof:             roof,
		Backyard:         backyard,
		SwimmingPool:     swimmingPool,
		HeatingSystem:    heatingSystem,
		Plumbing:         plumbing,
		ElectricalWiring: electricalWiring,
		PorchStyle:       porchStyle,
	}
}

func main() {
	// Chỉ cần nhà đơn giản nhưng vẫn phải truyền đủ tham số
	house := NewHouse(4, true, true, 2, true, false, false, false, false, false, "")
	fmt.Printf("House: %+v\n", house)
}
