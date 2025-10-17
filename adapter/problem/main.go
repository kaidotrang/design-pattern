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

func main() {
	orange := Orange{}

	MakeJuice(orange) // Nước cam

}

// type Coconut struct{}

// func (c Coconut) CrackAndDrain() string {
// 	return "Nước dừa"
// }
