package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/quetzyg/aoc2021/pkg/input"
)

type number struct {
	value  int
	marked bool
}

const (
	boardSize = 5
)

type board [boardSize][boardSize]number

func (b *board) bingo(pos int) bool {
	horizontal := 0
	vertical := 0

	for i := 0; i < boardSize; i++ {
		if b[i][pos].marked {
			vertical++
		}

		if b[pos][i].marked {
			horizontal++
		}
	}

	return horizontal == boardSize || vertical == boardSize
}

func (b *board) mark(numbers []int) (int, int) {
	for i, n := range numbers {
		for x := 0; x < boardSize; x++ {
			for y := 0; y < boardSize; y++ {
				if b[x][y].value == n {
					b[x][y].marked = true
				}

				if b[y][x].value == n {
					b[y][x].marked = true
				}
			}

			if b.bingo(x) {
				return i + 1, n
			}
		}
	}

	return math.MaxInt, 0
}

var (
	numbers []int
	rows    int
	boards  []*board
)

var parseFunc = func(line int, data string) error {
	if line == 0 {
		for i, v := range strings.Split(data, ",") {
			n, err := strconv.Atoi(v)

			if err != nil {
				return fmt.Errorf("input type mismatch on column %d, want integer, got %T", i, v)
			}

			numbers = append(numbers, n)
		}

		return nil
	}

	// Skip empty lines
	if data == "" {
		return nil
	}

	if rows%boardSize == 0 {
		boards = append(boards, new(board))
	}

	row := strings.Fields(data)

	if len(row) != boardSize {
		return fmt.Errorf("input length mismatch, want %d, got %d", boardSize, len(row))
	}

	current := len(boards) - 1

	for i, v := range row {
		n, err := strconv.Atoi(v)

		if err != nil {
			return fmt.Errorf("input type mismatch on column %d, want integer, got %T", i, v)
		}

		boards[current][rows%boardSize][i] = number{
			value:  n,
			marked: false,
		}
	}

	rows++

	return nil
}

func part12() (int, int) {
	var winner *board
	var loser *board

	min := math.MaxInt
	max := math.MinInt
	winLast := 0
	loseLast := 0

	for _, b := range boards {
		tries, last := b.mark(numbers)

		if tries < min {
			min = tries

			winner = b
			winLast = last
		}

		if tries > max {
			max = tries

			loser = b
			loseLast = last
		}
	}

	winSum := 0
	loseSum := 0

	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			if !winner[x][y].marked {
				winSum += winner[x][y].value
			}

			if !loser[x][y].marked {
				loseSum += loser[x][y].value
			}
		}
	}

	return winLast * winSum, loseLast * loseSum
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

	win, lose := part12()

	fmt.Printf("part #1: %d\n", win)
	fmt.Printf("part #2: %d\n", lose)
}
