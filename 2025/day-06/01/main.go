package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if _, err := os.Stat("input.txt"); err == nil {
		data, err := os.ReadFile("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		input = string(data)
	}

	sc := bufio.NewScanner(strings.NewReader(strings.Trim(input, "\n")))

	// Workbook
	wkb := make([][]string, 0, 1024)

	for sc.Scan() {
		wkb = append(wkb, strings.Fields(sc.Text()))
	}

	sum := 0
	xdim := len(wkb[0])
	for x := range xdim {
		col := column(wkb, x)
		n := len(col)
		op, list := col[n-1], col[:n-1]
		sum += apply(op, list)
	}

	fmt.Println(sum)
}

func num(s string) int {
	n := 0
	for _, v := range s {
		n = n*10 + int(v-'0')
	}
	return n
}

// Zero-based column
func column(wkb [][]string, x int) []string {
	n := len(wkb)
	col := make([]string, 0, len(wkb))
	for y := 0; y < n; y += 1 {
		log.Printf("(%d,%d): %q", x, y, wkb[y][x])
		col = append(col, wkb[y][x])
	}
	return col
}

func apply(op string, list []string) int {
	switch op {
	case "*":
		return mul(list)
	case "+":
		return add(list)
	default:
		panic(op + "not implemented")
	}
	return 0
}

func add(list []string) int {
	n := 0
	for _, v := range list {
		n += num(v)
	}
	return n
}

func mul(list []string) int {
	n := 1
	for _, v := range list {
		n *= num(v)
	}
	return n
}

var input = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`
