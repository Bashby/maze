package maze

import (
	"fmt"

	"github.com/fogleman/gg"
)

// Drawing draws a maze
type Drawing struct {
	Context   *gg.Context
	CellWidth int
}

func (d *Drawing) String() string {
	return fmt.Sprintf("<Drawing>")
}

// CreateDrawing create a drawing given dimensions and a cell size
func CreateDrawing(width, height, cellWidth int) Drawing {
	dc := gg.NewContext(width*cellWidth, height*cellWidth)

	// Fill background
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetRGBA(0.9, 0.9, 0.9, 1.0)
	dc.Fill()

	return Drawing{Context: dc, CellWidth: cellWidth}
}

// Save saves a drawing to disk given a filename
func (d Drawing) Save(filename string) {
	err := d.Context.SavePNG(filename)
	if err != nil {
		panic(err)
	}
}
