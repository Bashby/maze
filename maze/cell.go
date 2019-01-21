package maze

import (
	"fmt"

	"github.com/bashby/maze/util"
)

// CellNeighbor neighboring cell in a maze
type CellNeighbor struct {
	edge *Edge
	cell *Cell
}

func (n CellNeighbor) String() string {
	var str string

	if n.edge.boundary {
		str = fmt.Sprintf("<CellNeighbor edge:%v>", n.edge)
	} else {
		str = fmt.Sprintf(
			"<CellNeighbor edge:%v, cell.pos:%v cell.visited:%v>",
			n.edge,
			n.cell.pos,
			n.cell.visited)
	}

	return str
}

// Neighbors holds the neighboring cells of a cell
type Neighbors struct {
	north *CellNeighbor
	east  *CellNeighbor
	south *CellNeighbor
	west  *CellNeighbor
}

// Cell a point in a maze
type Cell struct {
	pos      util.Position
	neighbor Neighbors
	visited  bool
}

func (c Cell) String() string {
	return fmt.Sprintf(
		"<Cell %v,%v visited:%t\n\tn:%v\n\te:%v\n\ts:%v\n\tw:%v\n>",
		c.pos.X,
		c.pos.Y,
		c.visited,
		c.neighbor.north,
		c.neighbor.east,
		c.neighbor.south,
		c.neighbor.west)
}

// CreateCell creates a new cell at position {x,y}
func CreateCell(x, y int) *Cell {
	cell := Cell{
		pos: util.Position{X: x, Y: y},
	}

	return &cell
}

// Link populates the edges of the cell with other cells or edges
func (c *Cell) Link(cells [][]*Cell) {
	x := c.pos.X
	y := c.pos.Y
	xMax := len(cells[y]) - 1
	yMax := len(cells) - 1

	if y == 0 {
		c.neighbor.north = &CellNeighbor{edge: &Edge{closed: true, boundary: true}}
	} else {
		targetCell := cells[y-1][x]
		if targetCell.neighbor.south != nil {
			c.neighbor.north = &CellNeighbor{edge: targetCell.neighbor.south.edge, cell: targetCell}
		} else {
			c.neighbor.north = &CellNeighbor{edge: &Edge{closed: true}, cell: targetCell}
		}
	}

	if x == xMax {
		c.neighbor.east = &CellNeighbor{edge: &Edge{closed: true, boundary: true}}
	} else {
		targetCell := cells[y][x+1]
		if targetCell.neighbor.west != nil {
			c.neighbor.east = &CellNeighbor{edge: targetCell.neighbor.west.edge, cell: targetCell}
		} else {
			c.neighbor.east = &CellNeighbor{edge: &Edge{closed: true}, cell: targetCell}
		}
	}

	if y == yMax {
		c.neighbor.south = &CellNeighbor{edge: &Edge{closed: true, boundary: true}}
	} else {
		targetCell := cells[y+1][x]
		if targetCell.neighbor.north != nil {
			c.neighbor.south = &CellNeighbor{edge: targetCell.neighbor.north.edge, cell: targetCell}
		} else {
			c.neighbor.south = &CellNeighbor{edge: &Edge{closed: true}, cell: targetCell}
		}
	}

	if x == 0 {
		c.neighbor.west = &CellNeighbor{edge: &Edge{closed: true, boundary: true}}
	} else {
		targetCell := cells[y][x-1]
		if targetCell.neighbor.east != nil {
			c.neighbor.west = &CellNeighbor{edge: targetCell.neighbor.east.edge, cell: targetCell}
		} else {
			c.neighbor.west = &CellNeighbor{edge: &Edge{closed: true}, cell: targetCell}
		}
	}
}

// Draw draws a cell onto a drawing
func (c Cell) Draw(d Drawing) {
	x := float64(c.pos.X)
	y := float64(c.pos.Y)
	width := float64(d.CellWidth)
	dc := d.Context

	tl := util.PositionFloat{X: x * width, Y: y * width}
	tr := util.PositionFloat{X: tl.X + width, Y: tl.Y}
	bl := util.PositionFloat{X: tl.X, Y: tl.Y + width}
	br := util.PositionFloat{X: tl.X + width, Y: tl.Y + width}

	// Plot
	if c.neighbor.north.edge.closed {
		dc.DrawLine(tl.X, tl.Y, tr.X, tr.Y)
	}
	if c.neighbor.east.edge.closed {
		dc.DrawLine(tr.X, tr.Y, br.X, br.Y)
	}
	if c.neighbor.south.edge.closed {
		dc.DrawLine(br.X, br.Y, bl.X, bl.Y)
	}
	if c.neighbor.west.edge.closed {
		dc.DrawLine(bl.X, bl.Y, tl.X, tl.Y)
	}

	// Draw
	dc.SetLineWidth(2.0)
	dc.SetRGBA(0, 0, 0, 1)
	dc.Stroke()
}

// // GetRandomNeighbors returns a random edge of the cell; errors if DNE
// func (c *Cell) GetRandomNeighbors() (*Edge, error) {
// 	seed := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(seed)

// 	neighborIndex := r.Intn(4)

// 	var edge Edge
// 	switch neighborIndex {
// 	case 0:
// 		edge = c.north
// 	case 1:
// 		edge = c.east
// 	case 2:
// 		edge = c.south
// 	case 3:
// 		edge = c.west
// 	}

// 	var error error
// 	if !edge.boundary && edge.cell == nil {
// 		error = errors.New("edge does not exist")
// 	}

// 	return &edge, error
// }
