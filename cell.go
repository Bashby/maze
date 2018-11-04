package main

// Cell is a point in a maze
type Cell struct {
	pos   position
	north Edge
	east  Edge
	south Edge
	west  Edge
}

// Edge represents the edge between two cells in the maze
type Edge struct {
	cell     *Cell
	closed   bool
	boundary bool
}

// CreateCell creates a new Cell at position {x,y}
func CreateCell(x, y int) Cell {
	return Cell{
		pos: position{x, y},
	}
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
