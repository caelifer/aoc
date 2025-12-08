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
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read program from input if provided
	program = string(input)

	dial := &Dial{value: 50, mod: 100}
	if err := dial.ExecuteProgram(program); err != nil {
		log.Fatalf("failed to excecute sequence: %v", err)
	}
	fmt.Println(dial)
}

type Dial struct {
	value, pin, mod int
}

func (d *Dial) String() string {
	return fmt.Sprintf("Dial: %d of [0 - %d] (pin: %d)", d.value, d.mod-1, d.pin)
}

func (d *Dial) ExecuteProgram(program string) error {
	sc := bufio.NewScanner(strings.NewReader(program))

	for sc.Scan() {
		if t := sc.Text(); t != "" {
			if err := d.Rotate(t); err != nil {
				return err
			}
		}
	}
	return sc.Err()
}

func (d *Dial) Rotate(code string) error {
	dir, off, err := readCode(code)
	if err != nil {
		return err
	}
	// Enforce invariants by ignoring whole dial rotations:
	// * offset is always less than number of elements on the dial (mod > off)
	off %= d.mod

	// Rotation icon
	ri := '⟳'
	// Halndle left turn into right using ring equality: `value - off == value + mod - off`
	if dir == 'L' {
		off = d.mod - off
		ri = '⟲'
	}
	// Standard cyclic offset calculation
	oldValue := d.value
	d.value = (d.value + off) % d.mod

	// Update pin
	if d.value == 0 {
		d.pin += 1
	}

	val := strconv.Itoa(d.value)
	/*
	if val == "0" {
		val = " " + bold(val)
	}
	*/

	log.Printf("rotation: %2d %c %4s ⊜ %2s", oldValue, ri, code, val)
	return nil
}

func (d *Dial) Pin() int {
	return d.pin
}

func readCode(code string) (byte, int, error) {
	dir, offb := code[0], code[1:]
	off, err := strconv.Atoi(string(offb))
	return dir, off, err
}

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"

	Bold = "\033[1m"
)

func bold(txt string) string {
	return Bold + Red + txt + Reset
}

// Test program
var program = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

// vim: :ts=4:vts=4:sw=4:noexpandtab:ai:showmatch:
