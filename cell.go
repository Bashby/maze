package main

import (
	"fmt"
)

// Cell is a point in a maze
type Cell struct {
	pos     Position
	north   Edge
	east    Edge
	south   Edge
	west    Edge
	visited bool
}

// Edge represents the edge between two cells in the maze
type Edge struct {
	cell     *Cell
	closed   bool
	boundary bool
}

func (e Edge) String() string {
	return fmt.Sprintf("<Edge closed:%t boundary:%t>", e.closed, e.boundary)
}

// CreateCell creates a new cell at position {x,y}
func CreateCell(x, y int) Cell {
	return Cell{
		pos: Position{x, y},
	}
}

func (c Cell) String() string {
	return fmt.Sprintf("<Cell %v,%v visited:%t n:%v e:%v s:%v w:%v>", c.pos.x, c.pos.y, c.visited, c.north, c.east, c.south, c.west)
}

// LinkCell populates the edges of the cell
func (c *Cell) LinkCell(cells [][]Cell) {
	x := c.pos.x
	y := c.pos.y
	xMax := len(cells[y]) - 1
	yMax := len(cells) - 1

	//fmt.Printf("Handling x:%v (of %v) y:%v (of %v)\n", x, xMax, y, yMax)
	if y == 0 {
		c.north = Edge{closed: true, boundary: true}
	} else {
		c.north = Edge{cell: &cells[y-1][x]}
	}

	if x == xMax {
		c.east = Edge{closed: true, boundary: true}
	} else {
		c.east = Edge{cell: &cells[y][x+1]}
	}

	if y == yMax {
		c.south = Edge{closed: true, boundary: true}
	} else {
		c.south = Edge{cell: &cells[y+1][x]}
	}

	if x == 0 {
		c.west = Edge{closed: true, boundary: true}
	} else {
		c.west = Edge{cell: &cells[y][x-1]}
	}
}

// Draw draws a cell onto a drawing
func (c Cell) Draw(d *Drawing) {
	fmt.Printf("Drawing cell: %v\n", c)

	x := float64(c.pos.x)
	y := float64(c.pos.y)
	width := float64(d.cellWidth)
	dc := d.context

	tl := PositionFloat{x * width, y * width}
	tr := PositionFloat{tl.x + width, tl.y}
	bl := PositionFloat{tl.x, tl.y + width}
	br := PositionFloat{tl.x + width, tl.y + width}

	// Plot
	if c.north.closed {
		dc.DrawLine(tl.x, tl.y, tr.x, tr.y)
	}
	if c.east.closed {
		dc.DrawLine(tr.x, tr.y, br.x, br.y)
	}
	if c.south.closed {
		dc.DrawLine(br.x, br.y, bl.x, bl.y)
	}
	if c.west.closed {
		dc.DrawLine(bl.x, bl.y, tl.x, tl.y)
	}

	// Draw
	dc.SetLineWidth(2.0)
	dc.SetRGBA(0, 0, 0, 1)
	dc.Stroke()
}
