package main

import (
	"fmt"
	"strings"
)

func Echo(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response[:len(response)-1] // remove trailing newline
}

func Invert(records [][]string) string {
	// First initialize inverted array
	nRow, nCol := len(records), len(records[0])
	result := make([][]string, nCol)
	for i := 0; i < nCol; i++ {
		result[i] = make([]string, nRow)
	}

	// next copy the values inverting along the way
	for r, row := range records {
		for c, value := range row {
			result[c][r] = value
		}
	}

	// Finally use the echo function to convert 2D array to string
	return Echo(result)
}
