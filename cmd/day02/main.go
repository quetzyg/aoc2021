package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/quetzyg/aoc2021/pkg/input"
)

const (
	forward = "forward"
	down    = "down"
	up      = "up"
)

var (
	horizontal int
	depth1     int
	depth2     int
	aim        int
)

var parseFunc = func(line int, data string) error {
	fields := strings.Fields(data)

	if len(fields) != 2 {
		return fmt.Errorf("expected 2 fields per line, got %d", len(fields))
	}

	units, err := strconv.Atoi(fields[1])

	if err != nil {
		return err
	}

	switch fields[0] {
	case up:
		depth1 -= units
		aim -= units
	case down:
		depth1 += units
		aim += units
	case forward:
		horizontal += units
		depth2 += aim * units
	default:
		return fmt.Errorf("unknown command: %s", fields[0])
	}

	return nil
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

	fmt.Printf("part #1: %d\n", horizontal*depth1)
	fmt.Printf("part #2: %d\n", horizontal*depth2)
}
