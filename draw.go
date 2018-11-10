package main

import (
	"github.com/fogleman/gg"
)

// Drawing draws a maze
type Drawing struct {
	context   *gg.Context
	cells     [][]Cell
	cellWidth int
}

// NewDrawing creates a new drawing
func NewDrawing(maze Maze, cellWidth int) Drawing {
	y := len(maze.cells)
	var x int
	if y > 0 {
		x = len(maze.cells[0])
	}
	dc := gg.NewContext(x*cellWidth, y*cellWidth)

	// Fill background
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetRGBA(0.9, 0.9, 0.9, 1.0)
	dc.Fill()
	return Drawing{context: dc, cellWidth: cellWidth, cells: maze.cells}
}

// DrawCells draws cells onto the drawing
func (d *Drawing) DrawCells() {
	for i := range d.cells {
		for j := range d.cells[i] {
			d.cells[i][j].Draw(d)
		}
	}
}

// Save saves a drawing to disk
func (d Drawing) Save(filename string) {
	d.context.SavePNG(filename)
}
