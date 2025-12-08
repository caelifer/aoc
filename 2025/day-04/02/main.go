package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if _, err := os.Stat("input.txt"); err == nil {
		// Read from input file
		data, err := os.ReadFile("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		gridMap = string(data)
	}

	g := loadGridMap(strings.Trim(gridMap, "\n"))

	n := 0
	for {
		fmt.Print("\033[2J\033[H")
		fmt.Println(g.OptimizedMap())

		time.Sleep(50 * time.Millisecond)

		// Count accessible roles
		x := g.CountAccessibleRolls()
		if x == 0 {
			break
		}
		// Update roll count
		n += x
		// "Remove" accessible roles and create new map
		newMap := strings.ReplaceAll(g.OptimizedMap(), "x", ".")
		// Reload map
		g = loadGridMap(strings.Trim(newMap, "\n"))
	}
	fmt.Println(n)
}

type Cell struct {
	// coordinate
	x, y int
	busy bool
}

func (c *Cell) IsBusy() bool {
	return c.busy
}

type Grid struct {
	// Dimensions zero based
	x, y int
	data []byte
}

func (g *Grid) At(x, y int) *Cell {
	return &Cell{
		x:    x,
		y:    y,
		busy: g.data[y*g.y+x] != '.',
	}
}

func (g *Grid) Map() string {
	return g.drawMap(func(x, y int) byte {
		c := g.At(x, y)
		var sym byte = '.'
		if c.IsBusy() {
			sym = '@'
		}
		return sym
	})
}

func (g *Grid) OptimizedMap() string {
	return g.drawMap(func(x, y int) byte {
		c := g.At(x, y)
		var sym byte = '.'
		if c.IsBusy() {
			// if n := g.Neighbors(c.x, c.y); n < 4 {
			if g.Neighbors(c.x, c.y) < 4 {
				sym = 'x'
			} else {
				sym = '@'
			}
		}
		return sym
	})
}

func (g *Grid) drawMap(symGen func(x, y int) byte) string {
	var buf strings.Builder
	buf.Grow(g.x * g.y)

	// Draw optimized map
	for y := 0; y < g.y; y += 1 {
		for x := 0; x < g.x; x += 1 {
			buf.WriteByte(symGen(x, y))
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (g *Grid) Neighbors(x, y int) int {
	ns := clean(neighs(x, y), g.x, g.y)
	// log.Printf("collected neighbours for cell(%d, %d): %+v", x, y, ns)

	// Calculate number of busy spots
	n := 0
	for _, v := range ns {
		busy := State(g.At(v[0], v[1]).IsBusy())
		// log.Printf("cell(%d,%d) neighbour at (%d,%d): %v", x, y, v[0], v[1], busy)
		if busy {
			n += 1
		}
	}
	return n
}

func (g *Grid) CountAccessibleRolls() int {
	n := 0
	// Evaluate all rolls on the grid
	for y := 0; y < g.y; y += 1 {
		for x := 0; x < g.x; x += 1 {
			c := g.At(x, y)
			if c.IsBusy() && g.Neighbors(c.x, c.y) < 4 {
				n += 1
			}
		}
	}
	return n
}

func neighs(x, y int) [][2]int {
	return [][2]int{
		// bottom row
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},

		// mid row
		{x - 1, y},
		{x + 1, y},

		// top row
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
}

func clean(cs [][2]int, x, y int) [][2]int {
	// At least as big
	res := make([][2]int, 0, len(cs))

	for _, v := range cs {
		if (v[0] >= 0 && v[0] < x) && (v[1] >= 0 && v[1] < y) {
			res = append(res, v)
		}
	}
	return res
}

func loadGridMap(m string) *Grid {
	data := make([]byte, 0, len(m))
	y := 0

	for _, v := range []byte(m) {
		if v == '\n' {
			y += 1
			continue
		}
		data = append(data, v)
	}

	return &Grid{
		data: data,
		y:    y + 1,
		x:    len(data) / (y + 1),
	}
}

type State bool

func (s State) String() string {
	if s {
		return "busy"
	}
	return "free"
}

var gridMap = `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`
