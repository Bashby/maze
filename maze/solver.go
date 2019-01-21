package maze

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/bashby/maze/util"
	log "github.com/sirupsen/logrus"
)

//go:generate stringer -type=SolverType

// SolverType types of solvers for mazes
type SolverType int

// SolverType types of solvers for mazes
const (
	GrowingTreeNewest SolverType = iota + 1
	GrowingTreeRandom
	GrowingTree50Split
)

// Solver builds a maze using various rules
type Solver struct {
	solverType   SolverType
	cells        []*Cell
	pendingCells map[util.Position]*Cell
}

// CreateSolver Creates a solver struct
func CreateSolver(solverType SolverType, cells []*Cell) Solver {
	solver := Solver{
		solverType:   solverType,
		cells:        cells,
		pendingCells: make(map[util.Position]*Cell),
	}

	// Add initial cell to pending
	solver.pendingCells[cells[0].pos] = cells[0]

	return solver
}

// GenerateMaze Creates a maze using a solver type
func (s *Solver) GenerateMaze() {
	log.Info(fmt.Sprintf("Generating maze using '%v', starting at %v", s.solverType, s.cells[0].pos))
	// Let C be a list of cells, initially empty. Add one cell to C, at random.
	// Choose a cell from C, and carve a passage to any unvisited neighbor of that cell, adding that neighbor to C as well. If there are no unvisited neighbors, remove the cell from C.
	// Repeat #2 until C is empty.

	// for {
	// 	cell, ok := s. .selectCell();
	// 	if !ok {
	// 		break
	// 	}

	// 	randomNeighbors := cell.

	// Process cells in maze until all visited
	for {
		// Select
		cell, idx, ok := s.selectCell()

		// Stop, when no cells left
		if ok != nil {
			break
		}

		log.Debug("Selected ", cell)

		// Process
		s.processCell(cell, idx)
	}
}

func (s *Solver) processCell(cell *Cell, idx int) {
	// Visit cell
	cell.visited = true

	// Shuffle neighbors
	neighbors := []*CellNeighbor{cell.neighbor.north, cell.neighbor.east, cell.neighbor.south, cell.neighbor.west}
	for _, i := range rand.Perm(len(neighbors)) {
		neighbor := neighbors[i]

		log.Debug(cell.pos, " trying ", i, neighbor)

		// If unvisited, add to cells
		if !neighbor.edge.boundary && !neighbor.cell.visited {
			// If currently pending, skip
			_, alreadySeen := s.pendingCells[neighbor.cell.pos]
			if !alreadySeen {
				neighbor.edge.closed = false
				s.pendingCells[neighbor.cell.pos] = neighbor.cell
				s.cells = append(s.cells, neighbor.cell)
				log.Debug("Added ", neighbor.cell)
				return
			}
		}
	}

	// All neighbors visited, remove this cell from solver
	log.Debug("Removing", cell)
	s.cells = append(s.cells[:idx], s.cells[idx+1:]...)
	return
}

func (s *Solver) selectCell() (c *Cell, idx int, err error) {
	if len(s.cells) == 0 {
		return nil, 0, errors.New("no cells")
	}

	switch s.solverType {
	case GrowingTreeNewest:
		c, idx, err = s.selectCellNewest()
	case GrowingTreeRandom:
		c, idx, err = s.selectCellRandom()
	case GrowingTree50Split:
		// 50/50 split of Newest and random cell
		choice := rand.Intn(2)

		if choice == 0 {
			c, idx, err = s.selectCellNewest()
		} else {
			c, idx, err = s.selectCellRandom()
		}
	}

	return
}

func (s *Solver) selectCellNewest() (*Cell, int, error) {
	if len(s.cells) == 0 {
		return nil, 0, errors.New("empty cells")
	}

	// Return most recently added cell
	idx := len(s.cells) - 1
	return s.cells[idx], idx, nil
}

func (s *Solver) selectCellRandom() (*Cell, int, error) {
	if len(s.cells) == 0 {
		return nil, 0, errors.New("empty cells")
	}

	// Random cell
	idx := rand.Intn(len(s.cells))
	return s.cells[idx], idx, nil
}

// Add a random cell from the maze to the solver

// 	solver := Solver{}
// 	solver.cells = append(solver.cells, m.getRandomCell())

// 	// While solver.cells is not empty
// 	//   Get "newest cell" i.e. the cell at the end of the slice
// 	//   Choose a random unvisited neighbor, delete wall to it, add it to cells
// 	//   If no neighbors are unvisited, remove cell from cells, get next "newest" cell
// 	// Repeat until cells is empty.

// 	for len(solver.cells) > 0 {
// 		targetCell := solver.cells[len(solver.cells)-1]
// 		targetCell.visited = true

// 		fmt.Println(solver.cells)
// 		// Get random neighbor:
// 		seed := rand.NewSource(time.Now().UnixNano())
// 		r := rand.New(seed)
// 		neighbors := r.Perm(4)
// 		for i := range neighbors {
// 			neighbors[i]++
// 		}

// 		fmt.Println(neighbors)

// 		// for idx := range neighbors {
// 		// 	stop := false
// 		// 	switch idx {
// 		// 	case 1:
// 		// 		cell := processNeighbor(&targetCell.north)
// 		// 		if cell != nil {
// 		// 			solver.cells = append(solver.cells, cell)
// 		// 			stop = true
// 		// 			break
// 		// 		}
// 		// 		fallthrough
// 		// 	case 2:
// 		// 		cell := processNeighbor(&targetCell.east)
// 		// 		if cell != nil {
// 		// 			solver.cells = append(solver.cells, cell)
// 		// 			stop = true
// 		// 			break
// 		// 		}
// 		// 		fallthrough
// 		// 	case 3:
// 		// 		cell := processNeighbor(&targetCell.south)
// 		// 		if cell != nil {
// 		// 			solver.cells = append(solver.cells, cell)
// 		// 			stop = true
// 		// 			break
// 		// 		}
// 		// 		fallthrough
// 		// 	case 4:
// 		// 		cell := processNeighbor(&targetCell.west)
// 		// 		if cell != nil {
// 		// 			solver.cells = append(solver.cells, cell)
// 		// 			stop = true
// 		// 			break
// 		// 		}
// 		// 		fallthrough
// 		// 	default:
// 		// 		solver.cells = solver.cells[:len(solver.cells)-1]
// 		// 	}
// 		// 	if stop {
// 		// 		break
// 		// 	}
// 		// }

// 		// if !targetCell.north.boundary && !targetCell.north.cell.visited {
// 		// 	targetCell.north.closed = false
// 		// 	solver.cells = append(solver.cells, targetCell.north.cell)
// 		// } else if !targetCell.east.boundary && !targetCell.east.cell.visited {
// 		// 	targetCell.east.closed = false
// 		// 	solver.cells = append(solver.cells, targetCell.east.cell)
// 		// } else if !targetCell.south.boundary && !targetCell.south.cell.visited {
// 		// 	targetCell.south.closed = false
// 		// 	solver.cells = append(solver.cells, targetCell.south.cell)
// 		// } else if !targetCell.west.boundary && !targetCell.west.cell.visited {
// 		// 	targetCell.west.closed = false
// 		// 	solver.cells = append(solver.cells, targetCell.west.cell)
// 		// } else {
// 		// 	solver.cells = solver.cells[:len(solver.cells)-1]
// 		// }
// 	}

// 	//fmt.Println(m.cells)
// }

// func processNeighbor(edge *Edge) *Cell {
// 	if !edge.boundary && !edge.cell.visited {
// 		edge.closed = true
// 		return edge.cell
// 	}
// 	return nil
// }
