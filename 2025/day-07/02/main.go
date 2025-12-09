package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	}
	fmt.Println(g)
	fmt.Println("n:", n)
	fmt.Println("z:", g.CountTimelines())
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
	fmt.Println(string(r(g.x)))
	for y := 0; y < g.y; y += 1 {
		buf.WriteString(fmt.Sprintf("%3d ", y))
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

func (g *Grid) CountTimelines() int {
	// Find 'S' position
	for x := 0; x < g.x; x += 1 {
		if g.At(x, 0) == 'S' {
			return walk(g, x, 1)
		}
	}
	return 0
}

var memo map[int]int

func (g *Grid) LookUp(x, y int) int {
	if memo == nil {
		memo = make(map[int]int, g.x*g.y)
	}
	if v, ok := memo[y*g.x+x]; ok {
		return v
	}
	v := walk(g, x, y)
	memo[y*g.x+x] = v
	return v
}

func walk(g *Grid, x, y int) int {
	if x < 0 || y < 0 || x >= g.x || y >= g.y {
		return 0
	}
	v := g.At(x, y)
	switch v {
	case '|':
		// last row
		if y >= g.y-1 {
			return 1
		}
		return g.LookUp(x, y+1)
	case '^':
		return g.LookUp(x-1, y+1) + g.LookUp(x+1, y+1)

	default:
		return 0
	}
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

func r(n int) string {
	var buf strings.Builder
	buf.Grow(n*3 + 12)
	// 100s positions
	buf.WriteString("   ")
	for i := 0; i < n; i += 1 {
		buf.WriteByte('0' + byte(i/100))
	}
	buf.WriteByte('\n')
	// 10s positions
	buf.WriteString("   ")
	for i := 0; i < n; i += 1 {
		d := i % 100 / 10
		buf.WriteByte('0' + byte(d))
	}
	buf.WriteByte('\n')
	// once positions
	buf.WriteString("   ")
	for i := 0; i < n; i += 1 {
		buf.WriteByte('0' + byte(i%10))
	}
	return buf.String()
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
