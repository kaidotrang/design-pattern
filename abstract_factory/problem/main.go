package main

import "fmt"

// Chair interface
type Chair interface {
	Sit()
}

// Sofa interface
type Sofa interface {
	Lie()
}

// CoffeeTable interface
type CoffeeTable interface {
	Place()
}

// Modern furniture
type ModernChair struct{}

func (ModernChair) Sit() { fmt.Println("Sitting on a Modern Chair") }

type ModernSofa struct{}

func (ModernSofa) Lie() { fmt.Println("Lying on a Modern Sofa") }

type ModernCoffeeTable struct{}

func (ModernCoffeeTable) Place() { fmt.Println("Placing items on a Modern Coffee Table") }

// Victorian furniture
type VictorianChair struct{}

func (VictorianChair) Sit() { fmt.Println("Sitting on a Victorian Chair") }

type VictorianSofa struct{}

func (VictorianSofa) Lie() { fmt.Println("Lying on a Victorian Sofa") }

type VictorianCoffeeTable struct{}

func (VictorianCoffeeTable) Place() { fmt.Println("Placing items on a Victorian Coffee Table") }

// FurnitureSet: grouping manually
type FurnitureSet struct {
	Chair
	Sofa
	CoffeeTable
}

func main() {
	sets := []FurnitureSet{
		// Consistent: all Modern
		{
			Chair:       ModernChair{},
			Sofa:        ModernSofa{},
			CoffeeTable: ModernCoffeeTable{},
		},
		// Consistent: all Victorian
		{
			Chair:       VictorianChair{},
			Sofa:        VictorianSofa{},
			CoffeeTable: VictorianCoffeeTable{},
		},
		// Mixed styles: Modern Chair + Victorian Sofa
		{
			Chair:       ModernChair{},
			Sofa:        VictorianSofa{},
			CoffeeTable: ModernCoffeeTable{},
		},
	}

	for i, set := range sets {
		fmt.Printf("Furniture Set #%d:\n", i+1)
		set.Sit()
		set.Lie()
		set.Place()
		fmt.Println("---")
	}
}
