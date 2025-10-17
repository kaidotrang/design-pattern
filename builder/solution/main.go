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

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{
		house: House{
			Walls:   4,
			Floor:   true,
			Door:    true,
			Windows: 2,
			Roof:    true,
		},
	}
}

func (b *HouseBuilder) WithBackyard() *HouseBuilder {
	b.house.Backyard = true
	return b
}

func (b *HouseBuilder) WithSwimmingPool() *HouseBuilder {
	b.house.SwimmingPool = true
	return b
}

func (b *HouseBuilder) WithHeatingSystem() *HouseBuilder {
	b.house.HeatingSystem = true
	return b
}

func (b *HouseBuilder) WithPlumbing() *HouseBuilder {
	b.house.Plumbing = true
	return b
}

func (b *HouseBuilder) WithElectricalWiring() *HouseBuilder {
	b.house.ElectricalWiring = true
	return b
}

func (b *HouseBuilder) WithPorchStyle(style string) *HouseBuilder {
	b.house.PorchStyle = style
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}

func main() {
	// Nhà đơn giản
	simpleHouse := NewHouseBuilder().Build()
	fmt.Printf("Simple House: %+v\n", simpleHouse)

	// Nhà cao cấp
	luxuryHouse := NewHouseBuilder().
		WithBackyard().
		WithSwimmingPool().
		WithHeatingSystem().
		WithPlumbing().
		WithElectricalWiring().
		WithPorchStyle("Modern").
		Build()

	fmt.Printf("Luxury House: %+v\n", luxuryHouse)
}
