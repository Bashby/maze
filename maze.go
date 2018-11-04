package main

// Maze is a 2d grid of cells with possible walls between each cell
type Maze struct {
	width  int
	height int
	cells  [][]Cell
}

func (m Maze) linkCells() {
	for i := range m.cells {
		for j := range m.cells[i] {
			m.cells[i][j].LinkCell(m.cells)
		}
	}
}

// CreateMaze creates a maze of dimension {x,y}
func CreateMaze(x int, y int) Maze {
	cells := make([][]Cell, y)
	for i := range cells {
		cells[i] = make([]Cell, x)

		// init cells
		for j := range cells[i] {
			cells[i][j] = CreateCell(j, i)
		}
	}
	maze := Maze{x, y, cells}
	maze.linkCells()
	return maze
}
