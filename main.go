package main

import (
	"github.com/jinfantesc/draw/plain"
)

const (
	W = 2560.0
	H = 1080.0
)

func main() {
	dc := plain.DrawBasic(W, H)
	dc.SavePNG("basic.png")

	dc = plain.DrawSin(W, H)
	dc.SavePNG("sin.png")
}
