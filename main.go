package main

import (
	"strconv"
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

func DivideChain(c string) []string {
	lines := strings.Split(c, "\n")
	var digits []string

	for i, line := range lines {
		splitedLines := splitIntoArrays(line)

		for range splitedLines {
			if i == 0 {
				digits = append(digits, "")
			}
		}

		for j, splited := range splitedLines {
			if i < 3 {
				digits[j] = digits[j] + splited
			}
		}

	}

	return digits
}

func splitIntoArrays(line string) []string {
	var result []string

	for i := 0; i < len(line); i += 3 {
		end := i + 3
		if end > len(line) {
			end = len(line)
		}
		result = append(result, line[i:end])
	}
	return result
}

func ParseAccountNumber(c string) string {
	var accountNumber string
	numbers := DivideChain(c)
	for _, number := range numbers {
		accountNumber += OCRToString(number)
	}
	return accountNumber
}

func isValidAccount(an string) bool {
	numbers := strings.Split(an, "")
	acountNumberLength := len(numbers)

	if acountNumberLength != 9 {
		return false
	}

	var total int
	for _, number := range numbers {
		cn, _ := strconv.Atoi(number)
		total = total + acountNumberLength*cn
		acountNumberLength = acountNumberLength - 1
	}

	return total%11 == 0
}
