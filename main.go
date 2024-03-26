package main

import (
	"fmt"
	"strings"
)

func OCRToString(OCRNumber string) string {
	var num string
	switch OCRNumber {
	case " _ | ||_|":
		num = "0"
	case "     |  |":
		num = "1"
	case " _  _||_ ":
		num = "2"
	case " _  _| _|":
		num = "3"
	case "   |_|  |":
		num = "4"
	case " _ |_  _|":
		num = "5"
	case " _ |_ |_|":
		num = "6"
	case " _   |  |":
		num = "7"
	case " _ |_||_|":
		num = "8"
	case " _ |_| _|":
		num = "9"
	}
	return num
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
