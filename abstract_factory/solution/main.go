package main

import "fmt"

type Chair interface {
	Sit()
}

type Sofa interface {
	Lie()
}

type CoffeeTable interface {
	Place()
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
	CreateCoffeeTable() CoffeeTable
}

type ModernChair struct{}

func (ModernChair) Sit() { fmt.Println("Sitting on a Modern Chair") }

type ModernSofa struct{}

func (ModernSofa) Lie() { fmt.Println("Lying on a Modern Sofa") }

type ModernCoffeeTable struct{}

func (ModernCoffeeTable) Place() { fmt.Println("Placing items on a Modern Coffee Table") }

type ModernFactory struct{}

func (ModernFactory) CreateChair() Chair             { return ModernChair{} }
func (ModernFactory) CreateSofa() Sofa               { return ModernSofa{} }
func (ModernFactory) CreateCoffeeTable() CoffeeTable { return ModernCoffeeTable{} }

type VictorianChair struct{}

func (VictorianChair) Sit() { fmt.Println("Sitting on a Victorian Chair") }

type VictorianSofa struct{}

func (VictorianSofa) Lie() { fmt.Println("Lying on a Victorian Sofa") }

type VictorianCoffeeTable struct{}

func (VictorianCoffeeTable) Place() { fmt.Println("Placing items on a Victorian Coffee Table") }

type VictorianFactory struct{}

func (VictorianFactory) CreateChair() Chair             { return VictorianChair{} }
func (VictorianFactory) CreateSofa() Sofa               { return VictorianSofa{} }
func (VictorianFactory) CreateCoffeeTable() CoffeeTable { return VictorianCoffeeTable{} }

type FurnitureSet struct {
	Chair
	Sofa
	CoffeeTable
}

func CreateFurnitureSet(factory FurnitureFactory) FurnitureSet {
	return FurnitureSet{
		Chair:       factory.CreateChair(),
		Sofa:        factory.CreateSofa(),
		CoffeeTable: factory.CreateCoffeeTable(),
	}
}

func main() {
	factories := []FurnitureFactory{
		ModernFactory{},
		VictorianFactory{},
	}

	for i, factory := range factories {
		fmt.Printf("Furniture Set #%d:\n", i+1)
		set := CreateFurnitureSet(factory)
		set.Sit()
		set.Lie()
		set.Place()
		fmt.Println("---")
	}
}
