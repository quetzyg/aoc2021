package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/quetzyg/aoc2021/pkg/input"
)

const (
	zero = '0'
)

type bit struct {
	ones  int
	zeros int
}

var (
	ogr  = make(map[int]string) // O2 Generator Rating
	csr  = make(map[int]string) // CO2 Scrubber Rating
	bits []bit
)

var parseFunc = func(line int, data string) error {
	if bits == nil {
		bits = make([]bit, len(data))
	}

	if len(data) != len(bits) {
		return fmt.Errorf("input length mismatch, want %d, got %d", len(bits), len(data))
	}

	for i, bit := range data {
		if bit != '0' {
			bits[i].ones++
		} else {
			bits[i].zeros++
		}
	}

	ogr[line] = data
	csr[line] = data

	return nil
}

func part1(bits []bit) int {
	var (
		gama    int
		epsilon int
		exp     = float64(len(bits))
	)

	for _, b := range bits {
		exp--
		pow := math.Pow(2, exp)

		if b.ones > b.zeros {
			gama += int(pow)
		} else {
			epsilon += int(pow)
		}
	}

	return gama * epsilon
}

func part2(m map[int]string, callback func([]int, []int) []int) int64 {
	var (
		i int
		n int64
	)

	for len(m) != 1 {
		var (
			zeros []int
			ones  []int
		)

		for k, s := range m {
			if s[i] == zero {
				zeros = append(zeros, k)
			} else {
				ones = append(ones, k)
			}
		}

		for _, k := range callback(zeros, ones) {
			delete(m, k)
		}

		i++
	}

	for _, v := range m {
		n, _ = strconv.ParseInt(v, 2, 64)
	}

	return n
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing input path argument")
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}

	defer f.Close()

	err = input.Parse(f, bufio.ScanLines, parseFunc)

	if err != nil {
		log.Fatalf("error parsing input: %v", err)
	}

	fmt.Printf("part #1: %d\n", part1(bits))

	mc := part2(ogr, func(zeros, ones []int) []int {
		if len(ones) >= len(zeros) {
			return zeros
		}

		return ones
	})

	lc := part2(csr, func(zeros, ones []int) []int {
		if len(zeros) <= len(ones) {
			return ones
		}

		return zeros
	})

	fmt.Printf("part #2: %d\n", mc*lc)
}
