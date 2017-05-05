package main

import (
	"fmt"
)

//go:generate enums $GOFILE

// [@enums
// type colors int
// -Grau colors = (iota + 1)
// -Weiss
// -Gelb
// -Grün
// -Rot
// -Blau
// -Lila = Weiss
// @]

// [@enums
// type values float64
// -pi values = (iota + 1) * 3.14159265359
// -pi2
// -pi3
// -pi4
// -pi1 = pi
// -pi10 = 10 * pi
// @]

func main() {
	fmt.Println("List of colors constants:")
	for i, clr := range colors_Constants {
		fmt.Printf("colors_Constants[%v] = %v corresponds to the constant named \"%s\"\n", i, int(clr), clr)
	}
	fmt.Println("List of values constants:")
	for i, v := range values_Constants {
		fmt.Printf("values_Constants[%v] = %f corresponds to the constant named \"%s\"\n", i, v, v)
	}
}
