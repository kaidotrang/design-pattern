package main

import "fmt"

type Juicer interface {
	ExtractJuice() string
}

func MakeJuice(j Juicer) {
	fmt.Println(j.ExtractJuice())
}

type Orange struct{}

func (o Orange) ExtractJuice() string {
	return "Nước cam"
}

type Coconut struct{}

func (c Coconut) CrackAndDrain() string {
	return "Nước dừa"
}

type CoconutAdapter struct {
	coconut Coconut
}

func (a CoconutAdapter) ExtractJuice() string {
	return a.coconut.CrackAndDrain()
}

func main() {
	orange := Orange{}
	coconut := Coconut{}
	adapter := CoconutAdapter{coconut: coconut}

	MakeJuice(orange)  // Nước cam
	MakeJuice(adapter) // Nước dừa
}
