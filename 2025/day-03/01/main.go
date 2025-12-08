package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sum := 0
	sc := bufio.NewScanner(getInput())
	for sc.Scan() {
		if v := sc.Text(); v != "" {
			charges := convert(v)
			vv := maxJolts(charges)
			fmt.Printf("%q: %v\n", v, num(vv))
			sum += num(vv)
		}
	}
	fmt.Println(sum)
}

func convert(s string) []byte {
	res := []byte(s)
	for i := 0; i < len(res); i += 1 {
		res[i] -= '0'
	}
	return res
}

func maxJolts(charges []byte) []byte {
	n := len(charges)
	// exclude last element
	s1 := charges[:n-1]
	log.Printf("s1: %+v", s1)
	// find position of the larges in s1
	p1 := li(s1)

	// find position of the larges in rest
	s2 := charges[p1+1:]
	log.Printf("s2: %+v", s2)
	p2 := li(s2)

	return []byte{s1[p1], s2[p2]}
}

// index of the larges number in slice
func li(s []byte) int {
	res := 0
	max := s[res]
	log.Printf("new max: %d @ %d", max, res)
	for i, v := range s[1:] {
		if max < v {
			max, res = v, i+1
			log.Printf("new max: %d @ %d", max, res)
		}
	}
	return res
}

func num(r []byte) int {
	res := 0
	for _, v := range r {
		res = res*10 + int(v)
	}
	return res
}

func getInput() io.Reader {
	if _, err := os.Stat("input.txt"); err != nil {
		return strings.NewReader(joltages)
	}

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(data)
}

var joltages = `
987654321111111
811111111111119
234234234234278
818181911112111
`
