package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Maze is a 2d grid of cells with possible walls between each cell
type Maze struct {
	width  int
	height int
	cells  [][]Cell
}

// Solver builds a maze using various rules
type Solver struct {
	cells []*Cell
}

// CreateMaze creates a maze of dimension {x,y}
func CreateMaze(x int, y int) Maze {
	maze := Maze{width: x, height: y}
	maze.createCells()
	maze.linkCells()
	maze.placeWalls()

	return maze
}

func (m *Maze) createCells() {
	x := m.width
	y := m.height

	cells := make([][]Cell, y)
	for i := range cells {
		cells[i] = make([]Cell, x)

		// init cells
		for j := range cells[i] {
			cells[i][j] = CreateCell(j, i)
		}
	}

	m.cells = cells
}

func (m *Maze) linkCells() {
	for i := range m.cells {
		for j := range m.cells[i] {
			m.cells[i][j].LinkCell(m.cells)
		}
	}
}

func (m *Maze) getRandomCell() *Cell {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	randX := r.Intn(m.width)
	randY := r.Intn(m.height)

	return &m.cells[randY][randX]
}

func (m *Maze) placeWalls() {
	solver := Solver{}
	solver.cells = append(solver.cells, m.getRandomCell())

	// While solver.cells is not empty
	//   Get "newest cell" i.e. the cell at the end of the slice
	//   Choose a random unvisited neighbor, delete wall to it, add it to cells
	//   If no neighbors are unvisited, remove cell from cells, get next "newest" cell
	// Repeat until cells is empty.

	for len(solver.cells) > 0 {
		targetCell := solver.cells[len(solver.cells)-1]
		targetCell.visited = true

		// fmt.Println(solver.cells)
		// Get random neighbor:
		seed := rand.NewSource(time.Now().UnixNano())
		r := rand.New(seed)
		neighbors := r.Perm(4)
		for i := range neighbors {
			neighbors[i]++
		}

		fmt.Println(neighbors)

		for idx := range neighbors {
			stop := false
			switch idx {
			case 1:
				cell := processNeighbor(&targetCell.north)
				if cell != nil {
					solver.cells = append(solver.cells, cell)
					stop = true
					break
				}
				fallthrough
			case 2:
				cell := processNeighbor(&targetCell.east)
				if cell != nil {
					solver.cells = append(solver.cells, cell)
					stop = true
					break
				}
				fallthrough
			case 3:
				cell := processNeighbor(&targetCell.south)
				if cell != nil {
					solver.cells = append(solver.cells, cell)
					stop = true
					break
				}
				fallthrough
			case 4:
				cell := processNeighbor(&targetCell.west)
				if cell != nil {
					solver.cells = append(solver.cells, cell)
					stop = true
					break
				}
				fallthrough
			default:
				solver.cells = solver.cells[:len(solver.cells)-1]
			}
			if stop {
				break
			}
		}

		// if !targetCell.north.boundary && !targetCell.north.cell.visited {
		// 	targetCell.north.closed = false
		// 	solver.cells = append(solver.cells, targetCell.north.cell)
		// } else if !targetCell.east.boundary && !targetCell.east.cell.visited {
		// 	targetCell.east.closed = false
		// 	solver.cells = append(solver.cells, targetCell.east.cell)
		// } else if !targetCell.south.boundary && !targetCell.south.cell.visited {
		// 	targetCell.south.closed = false
		// 	solver.cells = append(solver.cells, targetCell.south.cell)
		// } else if !targetCell.west.boundary && !targetCell.west.cell.visited {
		// 	targetCell.west.closed = false
		// 	solver.cells = append(solver.cells, targetCell.west.cell)
		// } else {
		// 	solver.cells = solver.cells[:len(solver.cells)-1]
		// }
	}

	//fmt.Println(m.cells)
}

func processNeighbor(edge *Edge) *Cell {
	if !edge.boundary && !edge.cell.visited {
		edge.closed = true
		return edge.cell
	}
	return nil
}
