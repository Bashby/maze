package main

import "fmt"

type position struct {
	x int
	y int
}

func main() {
	var width, height int = 29, 23
	maze := CreateMaze(width, height)
	fmt.Println(maze)
}
