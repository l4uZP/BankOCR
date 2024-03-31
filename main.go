package main

import (
	"strconv"
	"strings"
)

func ParseAccountNumber(c string) string {
	var accountNumber string
	numbers := DivideChain(c)
	for _, number := range numbers {
		accountNumber += OCRToString(number)
	}
	return accountNumber
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

func OCRToString(OCRNumber string) string {
	number := numbers[OCRNumber]
	if number == "" {
		return "?"
	}
	return number
}

var numbers = map[string]string{
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

func isValidAccount(an string) bool {
	numbers := strings.Split(an, "")
	return hasCorrectLenght(len(numbers)) && isValidChecksum(numbers)
}

func hasCorrectLenght(accountNumberLength int) bool {
	return accountNumberLength == 9
}

func isValidChecksum(numbers []string) bool {
	var total int
	accountNumberLength := len(numbers)
	for _, number := range numbers {
		cn, _ := strconv.Atoi(number)
		total = total + accountNumberLength*cn
		accountNumberLength = accountNumberLength - 1
	}

	return total%11 == 0
}

func GetAccountsWithStatus(list []string) [][2]string {
	var accountsList [][2]string
	for _, account := range list {
		status := setAccountStatus(account)
		accountsList = append(accountsList, [2]string{0: account, 1: status})
	}

	return accountsList
}

func setAccountStatus(account string) string {
	status := ""
	if !isValidAccount(account) {
		status = "ERR"
	}

	if strings.Contains(account, "?") {
		status = "ILL"
	}
	return status
}
