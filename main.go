package main

import (
	"github.com/jinfantesc/draw/plain"
)

const (
	W = 512.0
	H = 512.0
)

func main() {
	dc := plain.DrawBasic(W, H)
	dc.SavePNG("out/basic.png")

	dc = plain.DrawSin(W, H)
	dc.SavePNG("out/sin.png")
}
