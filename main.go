package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Sequence []int
type Grid map[int][]int

var sequenceType string
var limit int

const (
	defaultSequenceType = "primes"
	defaultLimit        = 10
)

func init() {
	flag.StringVar(&sequenceType, "type", defaultSequenceType, "type of sequence (primes or fibonacci)")
	flag.IntVar(&limit, "limit", defaultLimit, "length of sequence")
	flag.Parse()
}

func main() {
	sequenceGenerator := NewSequenceGenerator(sequenceType, limit)
	sequence, err := sequenceGenerator.generate()
	if err != nil {
		log.Fatal(err)
	}

	gridGenerator := NewGridGenerator(sequence)
	grid, err := gridGenerator.generateGrid()
	if err != nil {
		log.Fatal(err)
	}

	formatter := NewFormatter(grid)
	formatter.formatGrid()
}

type SequenceGenerator struct {
	sequenceType string
	limit        int
}

func NewSequenceGenerator(sequenceType string, limit int) *SequenceGenerator {
	return &SequenceGenerator{sequenceType, limit}
}

func (sg *SequenceGenerator) generate() (Sequence, error) {
	switch sg.sequenceType {
	case "primes":
		return sg.generatePrimes(), nil
	case "fibonacci":
		return sg.generateFibonacci(), nil
	}

	return Sequence{}, errors.New("No associated sequence type.")
}

func (sg *SequenceGenerator) generatePrimes() Sequence {
	primes := Sequence{}
	number := 2

	for {
		if len(primes) == sg.limit {
			break
		}

		if sg.isAPrime(primes, number) == true {
			primes = append(primes, number)
		}

		number += 1
	}

	return primes
}

func (sg *SequenceGenerator) isAPrime(primes Sequence, number int) bool {
	for _, prime := range primes {
		if (number % prime) == 0 {
			return false
		}
	}

	return true
}

func (sg *SequenceGenerator) generateFibonacci() Sequence {
	seq := Sequence{1, 1}

	for i := 1; i < 5; i++ {
		seq = append(seq, sg.sumOfLastTwoElements(seq))
	}

	return seq
}

func (sg *SequenceGenerator) sumOfLastTwoElements(seq Sequence) int {
	length := len(seq)

	return seq[length-1] + seq[length-2]
}

type GridGenerator struct {
	sequence Sequence
}

func NewGridGenerator(sequence Sequence) *GridGenerator {
	return &GridGenerator{sequence}
}

func (gg *GridGenerator) generateGrid() (Grid, error) {
	grid := Grid{}

	// Append header row
	headerRow, err := gg.headerRow(gg.sequence)
	if err != nil {
		return nil, err
	}

	grid[0] = headerRow

	// Append all other rows
	for index, _ := range gg.sequence {
		otherRow, err := gg.otherRow(gg.sequence, index)
		if err != nil {
			return nil, err
		}

		grid[index+1] = otherRow
	}

	return grid, nil
}

// TODO: Return errors
func (gg *GridGenerator) headerRow(sequence Sequence) (Sequence, error) {
	row := append(sequence[:0:0], sequence...)
	row = append(Sequence{0}, row...)
	return row, nil
}

// TODO: Return errors
func (gg *GridGenerator) otherRow(sequence Sequence, rowIndex int) (Sequence, error) {
	row := Sequence{}

	for colIndex, _ := range sequence {
		if colIndex == 0 {
			row = append(row, sequence[rowIndex]) // Sidebar values
		}

		row = append(row, (sequence[colIndex] * sequence[rowIndex]))
	}

	return row, nil
}

const extraPadding = 2

type Formatter struct {
	Grid         Grid
	extraPadding int
}

func NewFormatter(grid Grid) *Formatter {
	return &Formatter{grid, extraPadding}
}

func (f *Formatter) formatGrid() {
	for i := 0; i < len(f.Grid); i++ {
		f.formatRow(i)
	}
}

func (f *Formatter) formatRow(rowIndex int) {
	if rowIndex == 1 {
		fmt.Println(strings.Repeat("-----", len(f.Grid)))
	}

	row := f.Grid[rowIndex]

	for colIndex, number := range row {
		if colIndex == 1 {
			fmt.Print("|")
		}

		f.formatCell(number, colIndex)
	}

	fmt.Print("\n")
}

func (f *Formatter) formatCell(number, colIndex int) {
	numberAsString := strconv.Itoa(number)
	whitespace := strings.Repeat(" ", f.colWidth(colIndex, len(numberAsString)))
	fmt.Print(numberAsString + whitespace)
}

func (f *Formatter) colWidth(colIndex, numberOfDigits int) int {
	return f.largestWidthForColumn(colIndex) - numberOfDigits + f.extraPadding
}

func (f *Formatter) largestWidthForColumn(colIndex int) int {
	lastRowIndex := len(f.Grid[colIndex]) - 1
	return len(strconv.Itoa(f.Grid[colIndex][lastRowIndex]))
}
