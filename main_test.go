package main

import (
	"strconv"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type BankOCRSuite struct{}

var _ = Suite(&BankOCRSuite{})

func (b *BankOCRSuite) Test_OCRToString_TransformChainsOfAisolatedSymbolsToItsCorrectNumber_Zero(c *C) {
	zero := " _ | ||_|"
	result := OCRToString(zero)
	c.Assert(result, Equals, "0")
}

func (b *BankOCRSuite) Test_OCRToString_TransformChainsOfAisolatedSymbolsToItsCorrectNumber_AllNumbers(c *C) {
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

func (b *BankOCRSuite) Test_ParseAccountNumber_WorksWithBasicExample(c *C) {
	an := "    _  _     _  _  _  _  _ \n  | _| _||_||_ |_   ||_||_|\n  ||_  _|  | _||_|  ||_| _|\n                                "

	result := ParseAccountNumber(an)
	c.Assert(result, Equals, "123456789")
}

func (b *BankOCRSuite) Test_ParseAccountNumber_TestWithAChainOfZeros(c *C) {
	an := " _  _  _  _  _  _  _  _  _ \n| || || || || || || || || |\n|_||_||_||_||_||_||_||_||_|\n                           "

	result := ParseAccountNumber(an)
	c.Assert(result, Equals, "000000000")
}

func (b *BankOCRSuite) Test_isValidAccount_WorksWithBasicExample(c *C) {
	an := "345882865"

	result := isValidAccount(an)
	c.Assert(result, Equals, true)
}

func (b *BankOCRSuite) Test_isValidAccount_WorksWithAnotherBasicExample(c *C) {
	an := "345832868"

	result := isValidAccount(an)
	c.Assert(result, Equals, true)
}

func (b *BankOCRSuite) Test_isValidAccount_ReturnsFalseWhenAccountNumberIsNotCorrect(c *C) {
	an := "345831868"

	result := isValidAccount(an)
	c.Assert(result, Equals, false)
}

func (b *BankOCRSuite) Test_isValidAccount_ReturnsFalseWhenAccountNumberHasIncorrectLenght(c *C) {
	an1 := ""
	an2 := "15882861"
	an3 := "3458828650"

	result1 := isValidAccount(an1)
	result2 := isValidAccount(an2)
	result3 := isValidAccount(an3)

	c.Assert(result1, Equals, false)
	c.Assert(result2, Equals, false)
	c.Assert(result3, Equals, false)
}

func (b *BankOCRSuite) Test_ParseAccountNumber_SetsQuestionMarkWhenTheNumberIsNotRecognized(c *C) {
	an := "    _  _  _  _  _  _     _ \n|_||_|| || ||_   |  |  | _ \n  | _||_||_||_|  |  |  | _|\n                           "

	result := ParseAccountNumber(an)
	c.Assert(result, Equals, "49006771?")
}

func (b *BankOCRSuite) Test_GetAccountsWithStatus_ReturnsAListOfParsedAccountNumbersWithItsStatus(c *C) {
	accountNumbers := []string{"457508000", "664371495", "86110??36"}

	result := GetAccountsWithStatus(accountNumbers)
	c.Assert(result, DeepEquals, [][2]string{{"457508000", ""}, {"664371495", "ERR"}, {"86110??36", "ILL"}})
}

func (b *BankOCRSuite) Test_ParseScannedFileToAccountsList_CompleteTheWholeFlow(c *C) {
	scannedFile := []string{
		" _  _  _  _  _  _  _  _    \n| || || || || || || ||_   |\n|_||_||_||_||_||_||_| _|  |\n                           ",
		"    _  _  _  _  _  _     _ \n|_||_|| || ||_   |  |  | _ \n  | _||_||_||_|  |  |  | _|\n                           ",
		"    _  _     _  _  _  _  _ \n  | _| _||_| _ |_   ||_||_|\n  ||_  _|  | _||_|  ||_| _ \n                           ",
	}

	results := ParseScannedFileToAccountsList(scannedFile)
	c.Assert(results, DeepEquals, [][2]string{{"000000051", ""}, {"49006771?", "ILL"}, {"1234?678?", "ILL"}})
}
