package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	_, err := os.Stat("input.txt")
	if err == nil {
		data, err := os.ReadFile("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		input = string(data)
	}
	g := loadGrid(strings.Trim(input, "\n"))
	// fmt.Println(g)

	// Number of splits
	n := 0
	for y := 0; y < g.y; y += 1 {
		// Clear screen
		fmt.Print("\033[2J\033[H")
		for x := 0; x < g.x; x += 1 {
			v := g.At(x, y-1)
			switch v {
			case '@':
				continue
			case 'S':
				g.Write(x, y, '|')
			case '|':
				vv := g.At(x, y)
				if vv == '^' {
					// split the stream
					n += 1
					g.Write(x-1, y, '|')
					g.Write(x+1, y, '|')
				} else {
					g.Write(x, y, '|')
				}
			}
		}
		fmt.Println(g)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("n:", n)
}

type Grid struct {
	// Bounds
	x, y int
	// Data
	data []byte
}

func (g *Grid) String() string {
	var buf strings.Builder
	buf.Grow(g.x * g.y)
	for y := 0; y < g.y; y += 1 {
		for x := 0; x < g.x; x += 1 {
			buf.WriteByte(g.At(x, y))
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (g *Grid) At(x, y int) byte {
	if x < 0 || y < 0 || x >= g.x || y >= g.y {
		return '@' // out-of-bounds marker
	}
	return g.data[y*g.x+x]
}

func (g *Grid) Write(x, y int, v byte) {
	if x < 0 || y < 0 || x >= g.x || y >= g.y {
		return
	}
	g.data[y*g.x+x] = v
}

func loadGrid(m string) *Grid {
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

var input = `
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`
