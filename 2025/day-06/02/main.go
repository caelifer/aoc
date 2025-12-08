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
	wkb := make([][]rune, 0, 1024)

	for sc.Scan() {
		wkb = append(wkb, []rune(sc.Text()))
	}

	sum := 0
	xdim := len(wkb[0])

	nums := make([]string, 0, 4)
	for x := xdim - 1; x >= 0; x -= 1 {
		col := column(wkb, x)
		// Check for field border
		if isBlank(col) {
			continue
		}
		num, op := split(col)
		nums = append(nums, strings.Trim(string(num), " "))
		if op != ' ' {
			// Do calculation
			s := apply(string(op), nums)
			log.Printf("applying %c to %+v = %d", op, nums, s)
			sum += s
			// Reset nums
			nums = nums[:0]
		}
	}

	fmt.Println(sum)
}

func isBlank(rs []rune) bool {
	for _, r := range rs {
		if r != ' ' {
			return false
		}
	}
	return true
}

func split(rs []rune) ([]rune, rune) {
	n := len(rs)
	return rs[:n-1], rs[n-1]
}

func digit(r rune) int {
	return int(r - '0')
}

func num(s string) int {
	n := 0
	for _, v := range s {
		n = n*10 + digit(v)
	}
	return n
}

// Zero-based column
func column(wkb [][]rune, x int) []rune {
	n := len(wkb)
	col := make([]rune, 0, len(wkb))
	for y := 0; y < n; y += 1 {
		// log.Printf("(%d,%d): %q", x, y, wkb[y][x])
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
