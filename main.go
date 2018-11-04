package main

import "fmt"

// Position is a point in a 2d coordinate system
type Position struct {
	x int
	y int
}

// PositionFloat is a point in a 2d coordinate system, using floats
type PositionFloat struct {
	x float64
	y float64
}

func main() {
	var width, height int = 29, 23
	maze := CreateMaze(width, height)
	drawing := NewDrawing(maze, 20)
	drawing.DrawCells()
	drawing.Save("maze.png")

	fmt.Println("Done!")
}
