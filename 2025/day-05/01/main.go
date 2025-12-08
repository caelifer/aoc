package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	db := NewInventoryWithCapacity(1 << 10)
	spoiled := make([]string, 0, 1024)

	sc := bufio.NewScanner(strings.NewReader(strings.Trim(input, "\n")))
	// Scan ranges
	for sc.Scan() {
		ln := sc.Text()
		if ln == "" {
			log.Printf("done processing ranges")
			break
		}
		// Add range to DB
		db.Add(NewRange(ln))
		log.Printf("processed: %q - %v", ln, db)
	}
	// Test ingredient against DB
	for sc.Scan() {
		ln := sc.Text()
		if ln == "" {
			break
		}

		if !db.Test(ln) {
			// Found spoiled product
			spoiled = append(spoiled, ln)
			log.Printf("spoiled: %q", ln)
		}
	}

	// Report
	fmt.Println(len(spoiled))
}

func toNum(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("invalid number: %q: %v", s, err)
	}
	return n
}

type Range struct {
	lo, hi int
}

func NewRange(s string) *Range {
	ps := strings.SplitN(s, "-", 2)
	return &Range{
		lo: toNum(ps[0]),
		hi: toNum(ps[1]),
	}
}

func (r *Range) String() string {
	return fmt.Sprintf("(%d,%d)", r.lo, r.hi)
}

func (r *Range) IsInRange(s string) bool {
	x := toNum(s)
	// x is in range if x >= lo and x <= hi
	return x >= r.lo && x <= r.hi
}

type DB struct {
	ranges []*Range
}

func NewInventoryWithCapacity(cap int) *DB {
	return &DB{
		ranges: make([]*Range, 0, cap),
	}
}

func (db *DB) String() string {
	return fmt.Sprintf("{ ranges(%d): %+v }", len(db.ranges), db.ranges)
}

func (db *DB) Add(r *Range) {
	db.ranges = append(db.ranges, r)
}

func (db *DB) Test(s string) bool {
	for _, r := range db.ranges {
		if r.IsInRange(s) {
			return true
		}
	}
	return false
}

var input = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`
