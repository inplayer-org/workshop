package main

import (
	"repo.inplayer.com/workshop/test/pkg/geometry"
	"repo.inplayer.com/workshop/test/pkg/hello"
)

func main() {
	r := &geometry.Rect{Width: 10, Height: 20}
	c := &geometry.Circle{123}

	hello.Hello(r)
	hello.Hello(c)
}
