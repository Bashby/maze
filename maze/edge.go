package maze

import "fmt"

// Edge represents the edge between two cells in the maze
type Edge struct {
	closed   bool
	boundary bool
}

func (e Edge) String() string {
	return fmt.Sprintf("<Edge closed:%t boundary:%t>", e.closed, e.boundary)
}
