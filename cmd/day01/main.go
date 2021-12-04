package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/quetzyg/aoc2021/pkg/input"
)

const windowSize = 3

var measurements []int

var parseFunc = func(line int, data string) error {
	depth, err := strconv.Atoi(data)

	if err != nil {
		return err
	}

	measurements = append(measurements, depth)

	return nil
}

func part1(data []int) int {
	var (
		prev  int
		total int
	)

	for _, depth := range data {
		if prev == 0 {
			prev = depth

			continue
		}

		if depth > prev {
			total++
		}

		prev = depth
	}

	return total
}

func part2(data []int, window int) int {
	var (
		prev  int
		total int
	)

	for i := 0; i <= len(data)-window; i++ {
		sum := 0

		for _, depth := range data[i : i+window] {
			sum += depth
		}

		if prev == 0 {
			prev = sum

			continue
		}

		if sum > prev {
			total++
		}

		prev = sum
	}

	return total
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
		log.Fatalf("unable to parse input: %v", err)
	}

	fmt.Printf("part #1: %d\n", part1(measurements))
	fmt.Printf("part #2: %d\n", part2(measurements, windowSize))
}
