package main

import (
	"strconv"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BankOCRSuite struct{}

var _ = Suite(&BankOCRSuite{})

func (b *BankOCRSuite) Test_OCRToString_TransformChainsOfAisolatedSymbolsToItsCorrectNumberZero(c *C) {
	zero := " _ | ||_|"
	result := OCRToString(zero)
	c.Assert(result, Equals, "0")
}

func (b *BankOCRSuite) Test_OCRToString_TransformChainsOfAisolatedSymbolsToItsCorrectNumberAllNumbers(c *C) {
	numbers := []string{"     |  |", " _  _||_ ", " _  _| _|", "   |_|  |", " _ |_  _|", " _ |_ |_|", " _   |  |", " _ |_||_|", " _ |_| _|"}

	for i, number := range numbers {
		cn := strconv.Itoa(i + 1)
		result := OCRToString(number)
		c.Assert(result, DeepEquals, cn)
	}
}

func (b *BankOCRSuite) Test_DivideChain_ExtractTheNecessaryElementsForEachNumber(c *C) {
	an := " _  _  _  _  _  _  _  _  _ \n| || || || || || || || || |\n|_||_||_||_||_||_||_||_||_|\n                           "

	result := DivideChain(an)
	c.Assert(result, DeepEquals, []string{
		" _ | ||_|",
		" _ | ||_|",
		" _ | ||_|",
		" _ | ||_|",
		" _ | ||_|",
		" _ | ||_|",
		" _ | ||_|",
		" _ | ||_|",
		" _ | ||_|",
	})
}

func (b *BankOCRSuite) Test_DivideChain_ExtractTheNecessaryElementsForOtherNumbers(c *C) {
	an := "    _  _     _  _  _  _  _ \n  | _| _||_||_ |_   ||_||_|\n  ||_  _|  | _||_|  ||_| _|\n                                "

	result := DivideChain(an)
	c.Assert(result, DeepEquals, []string{
		"     |  |",
		" _  _||_ ",
		" _  _| _|",
		"   |_|  |",
		" _ |_  _|",
		" _ |_ |_|",
		" _   |  |",
		" _ |_||_|",
		" _ |_| _|",
	})
}
