package main

import (
	"fmt"
	"strconv"
	"strings"
)

const extraPadding = 2

type Series []int
type Grid map[int][]int

func main() {
	primes := generatePrimes(10)
	grid := generateGrid(primes)
	formatGrid(grid)
}

func generatePrimes(limit int) Series {
	primes := Series{}
	number := 2

	for {
		if len(primes) == limit {
			break
		}

		if isAPrime(primes, number) == true {
			primes = append(primes, number)
		}

		number += 1
	}

	return primes
}

func isAPrime(primes Series, number int) bool {
	for _, prime := range primes {
		if (number % prime) == 0 {
			return false
		}
	}

	return true
}

func generateGrid(series Series) Grid {
	grid := Grid{}

	// Append header row
	grid[0] = headerRow(series)

	// Append all other rows
	for index, _ := range series {
		grid[index+1] = otherRow(series, index)
	}

	return grid
}

func headerRow(series Series) Series {
	row := append(series[:0:0], series...)
	row = append(Series{0}, row...)
	return row
}

func otherRow(series Series, rowIndex int) Series {
	row := Series{}

	for colIndex, _ := range series {
		if colIndex == 0 {
			row = append(row, series[rowIndex]) // Sidebar values
		}

		row = append(row, (series[colIndex] * series[rowIndex]))
	}

	return row
}

func formatGrid(grid Grid) {
	for i := 0; i < len(grid); i++ {
		formatRow(grid, i)
	}
}

func formatRow(grid Grid, rowIndex int) {
	if rowIndex == 1 {
		fmt.Println(strings.Repeat("-----", len(grid)))
	}

	row := grid[rowIndex]

	for colIndex, number := range row {
		if colIndex == 1 {
			fmt.Print("|")
		}

		formatCell(grid, number, colIndex)
	}

	fmt.Print("\n")
}

func formatCell(grid Grid, number, colIndex int) {
	numberAsString := strconv.Itoa(number)
	whitespace := strings.Repeat(" ", colWidth(grid, colIndex, len(numberAsString)))
	fmt.Print(numberAsString + whitespace)
}

func colWidth(grid Grid, colIndex, numberOfDigits int) int {
	return largestWidthForColumn(grid, colIndex) - numberOfDigits + extraPadding
}

func largestWidthForColumn(grid Grid, colIndex int) int {
	lastRowIndex := len(grid[colIndex]) - 1
	return len(strconv.Itoa(grid[colIndex][lastRowIndex]))
}
