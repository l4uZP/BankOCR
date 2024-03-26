package main

import (
	"fmt"
	"strings"
)

func OCRToString(OCRNumber string) string {
	numbers := map[string]string{
		" _ | ||_|": "0",
		"     |  |": "1",
		" _  _||_ ": "2",
		" _  _| _|": "3",
		"   |_|  |": "4",
		" _ |_  _|": "5",
		" _ |_ |_|": "6",
		" _   |  |": "7",
		" _ |_||_|": "8",
		" _ |_| _|": "9",
	}
	return numbers[OCRNumber]
}

func DivideChain(c string) [][]string {
	lines := strings.Split(c, "\n")
	var digits [][]string

	for i, _ := range digits {
		digits[i] = make([]string, len(lines[0])/3)
		fmt.Println(digits[i])
	}

	return digits
}
