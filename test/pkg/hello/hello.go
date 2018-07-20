package hello

import (
	"fmt"

	"repo.inplayer.com/workshop/test/pkg/geometry"
)

func Hello(shape geometry.Shape) {
	fmt.Printf("Sssehape area: %f\n", shape.Area())
	fmt.Printf("Shape perimeter: %f\n", shape.Perim())
}
