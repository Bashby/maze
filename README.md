<p align="center">
  <img src="static\logo.png" width="750" height="150" title="Mazeclient Logo" style="border-radius: 5px;">
</p>

# maze

Maze Generation using GoLang

## Getting started

### Dependencies

You need [dep](https://github.com/golang/dep) for dependency management.

    dep ensure

### Building

    go build -o mazegen.exe

### Interacting

You can control the seed, and dimensions of mazes you generate.

The generated maze will be in the same folder under the name `maze.png`.

Use `-h` to get help on using the generator.

    mazegen.exe -h

Example output:

```
Usage of ...\maze\mazegen.exe:
  -debug
        debug mode; high verbosity
  -height int
        maze height, in cells (default 5)
  -seed int
        pseudo RNG seed (default 1553884783481282100)
  -width int
        maze width, in cells (default 5)
```

#### Examples

Create a 5 by 5 maze:

    mazegen.exe -height 5 -width -5
