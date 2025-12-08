package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	// Print merged ranges
	fmt.Printf("%v\n", db)

	// Print fresh count
	fmt.Printf("%v\n", db.CountFresh())
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
	slices.SortFunc(db.ranges, func(a, b *Range) int {
		return a.lo - b.lo
	})
	db.Merge()
}

func (db *DB) Test(s string) bool {
	for _, r := range db.ranges {
		if r.IsInRange(s) {
			return true
		}
	}
	return false
}

func (db *DB) CountFresh() int {
	n := 0
	for _, r := range db.ranges {
		n += r.hi - r.lo + 1
	}
	return n
}

func (db *DB) Merge() {

	if len(db.ranges) < 2 {
		return
	}

	slices.SortFunc(db.ranges, func(a, b *Range) int {
		return a.lo - b.lo
	})

	ranges := make([]*Range, 0, len(db.ranges))
	for _, r := range db.ranges {
		if len(ranges) == 0 {
			ranges = append(ranges, r)
			continue
		}
		// get the last element
		r1 := ranges[len(ranges)-1]
		if r1.hi >= r.lo-1 {
			r1 = merge(r1, r)
			ranges[len(ranges)-1] = r1
		} else {
			ranges = append(ranges, r)
		}
	}
	db.ranges = ranges
}

func merge(r1, r2 *Range) *Range {
	return &Range{
		lo: r1.lo,
		hi: max(r1.hi, r2.hi),
	}
}

var input = `
3-5
10-14
16-20
12-18
100000004-100000004
100000000-100000001
100000000-100000009

1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
`
