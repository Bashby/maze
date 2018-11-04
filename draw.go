package main

import "github.com/fogleman/gg"

func draw(width, height int) {
	dc := gg.NewContext(1000, 1000)
	dc.SavePNG("out.png")
}
