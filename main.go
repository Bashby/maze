package main

import (
	"flag"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bashby/maze/maze"
)

func main() {
	log.Info("Starting...")

	// Arguments
	widthPtr := flag.Int("width", 5, "maze width, in cells")
	heightPtr := flag.Int("height", 5, "maze height, in cells")
	debugPtr := flag.Bool("debug", false, "debug mode; high verbosity")
	randSeedPtr := flag.Int64("seed", time.Now().UTC().UnixNano(), "pseudo RNG seed")
	flag.Parse()

	// Seed pRNG
	rand.Seed(*randSeedPtr)

	// Handle debugging
	if *debugPtr {
		log.SetLevel(log.DebugLevel)
	}

	// Generate maze
	maze := maze.Create(*widthPtr, *heightPtr, maze.GrowingTree50Split)
	log.Debug(maze)

	// Render maze
	maze.Save("maze.png")

	log.Info("Done!")
}
