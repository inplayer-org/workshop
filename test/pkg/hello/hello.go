package hello

import (
	"fmt"

	"repo.inplayer.com/workshop/test/pkg/geometry"
)

func Hello(shape geometry.Shape) {
<<<<<<< HEAD
	fmt.Printf("asdad area: %f\n", shape.Area())
=======
	fmt.Printf("Sssehape area: %f\n", shape.Area())
>>>>>>> d976ea68801608481c7c14cc5403136a55f95cde
	fmt.Printf("Shape perimeter: %f\n", shape.Perim())
}
