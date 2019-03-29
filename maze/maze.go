package maze

import (
	"fmt"
	"math/rand"
)

// Maze is a 2d grid of cells with possible walls between each cell
type Maze struct {
	width  int
	height int
	cells  [][]*Cell
}

func (m Maze) String() string {
	var cellsStr string
	for _, i := range m.cells {
		for _, j := range i {
			cellsStr += fmt.Sprintf("%v\n", j)
		}
	}
	return fmt.Sprintf(
		"<Maze %v,%v cells:\n%v>",
		m.width,
		m.height,
		cellsStr)
}

// Create creates a maze of dimension {width, height}
func Create(width int, height int, solverType SolverType) Maze {
	maze := Maze{width: width, height: height}
	maze.createCells()
	maze.linkCells()
	maze.placeWalls(solverType)

	return maze
}

// Save draw the maze and save to the provided filename
func (m *Maze) Save(fileName string) {
	drawing := CreateDrawing(m.width, m.height, 15)

	for i := range m.cells {
		for j := range m.cells[i] {
			m.cells[i][j].Draw(drawing)
		}
	}

	drawing.Save(fileName)
}

func (m *Maze) createCells() {
	x := m.width
	y := m.height

	cells := make([][]*Cell, y)
	for i := range cells {
		cells[i] = make([]*Cell, x)

		for j := range cells[i] {
			cells[i][j] = CreateCell(j, i)
		}
	}

	m.cells = cells
}

func (m *Maze) linkCells() {
	for _, row := range m.cells {
		for _, cell := range row {
			cell.Link(m.cells)
		}
	}
}

func (m *Maze) getRandomCell() *Cell {
	randX := rand.Intn(m.width)
	randY := rand.Intn(m.height)

	return m.cells[randY][randX]
}

func (m *Maze) placeWalls(solverType SolverType) {
	randCell := m.getRandomCell()
	cells := []*Cell{randCell}
	solver := CreateSolver(solverType, cells)
	solver.GenerateMaze()
}
