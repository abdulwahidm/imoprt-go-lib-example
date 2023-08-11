package main

import (
	"fmt"

	"github.com/abdulwahidm/geometry-lib/shape"
)

func main() {
	firstFloor := shape.Rectangle{Width: 7.5, Length: 6.5}
	secondFloor := shape.Rectangle{Width: 4.5, Length: 5.5}

	// fmt.Println("Luas area lantai pertama adalah:", firstFloor.Area())
	// fmt.Println("Luas area lantai kedua adalah:", secondFloor.Area())

	totalArea := firstFloor.Area() + secondFloor.Area()
	fmt.Println("Total Area: ", totalArea, "m2")
}
