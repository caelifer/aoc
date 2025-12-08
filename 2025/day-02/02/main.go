package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// scanner
	sc := bufio.NewScanner(getInput())
	// split input on ','
	sc.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		ln := len(data)
		if ln == 0 {
			return 0, nil, bufio.ErrFinalToken
		}
		if n := bytes.Index(data, []byte{','}); n >= 0 {
			return n + 1, data[:n], nil
		}
		return ln, data, bufio.ErrFinalToken
	})

	// final sum of all found invalid IDs
	sum := 0

	// scan input
	for sc.Scan() {
		// scanned token
		if v := strings.TrimRight(sc.Text(), "\n"); v != "" {
			log.Printf("token: %q", v)

			// create new ID range
			r, err := NewRange(v)
			if err != nil {
				log.Fatalf("bad id range: %q: %v", v, err)
			}
			// sum all invalid IDs in a range
			for _, id := range r.InvalidIDs() {
				sum += id.Value()
			}
		}
	}
	fmt.Printf("sum: %d, match checksum %d: %v\n", sum, checkSum, sum == checkSum)
}

// Range type to represent ID range.
type Range struct {
	pat string
	// inclusive bounds
	lo, hi int
}

// NewRange construct new ID range object.
func NewRange(pat string) (*Range, error) {
	// Split on `-`
	parts := strings.Split(pat, "-")
	// Parse lower bound
	lo, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	// Parse upper bound
	hi, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return &Range{
		pat: pat,
		lo:  lo,
		hi:  hi,
	}, nil
}

// InvalidIDs generates a list of invalid IDs found in a range.
func (r *Range) InvalidIDs() []*ID {
	res := make([]*ID, 0, 1024)
	for i := r.lo; i <= r.hi; i += 1 {
		id := NewID(i)
		if !id.IsValid() {
			res = append(res, id)
		}
	}
	return res
}

// ID type represents ID
type ID struct {
	id  string
	val int
}

// NewID creates an instance of ID based on then integer value.
func NewID(val int) *ID {
	return &ID{id: strconv.Itoa(val), val: val}
}

// Value returns ID's integer value
func (id *ID) Value() int {
	return id.val
}

// IsValid checks if ID is valid.
func (id *ID) IsValid() bool {
	n := len(id.id)
	for i := n / 2; i > 0; i -= 1 {
		if match(id.id, id.id[:i], n/i) {
			log.Printf("bad id: %s - %s repeats %d times", id.id, id.id[:i], n/i)
			return false
		}
	}
	return true
}

// String returns ID's string value.
func (id *ID) String() string {
	return id.id
}

// Check if `pat` fits `tgt` exactly `n` times.
func match(tgt, pat string, n int) bool {
	nt, np := len(tgt), len(pat)
	if np*n != nt {
		// Do not fit whole times - negative match
		return false
	}

	// Check if pattern fits repeatedly starting from the `np` offset
	for i := np; i < nt; i += np {
		if tgt[i:i+np] != pat {
			// Return negative match
			return false
		}
	}
	// Repeating pattern fits the target
	return true
}

func getInput() io.Reader {
	if _, err := os.Stat("input.txt"); err != nil {
		return strings.NewReader(input)
	}
	// Read file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(data)
}

// Test input
var input = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

// Test checksum
var checkSum = 34421651192
