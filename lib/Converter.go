/**
 * Eric Schmidt
 * CSS 508 Spring 2023
 * 2023-05-13
 *
 * The utility class that includes the `IntegerToWordedString` as a public
 * function. (Golang allows functions outside of structs.)
 *
 */
package lib

import (
	"fmt"
	"strings"
)

func IntegerToWordedString(number uint) (string, error) {

	// Get the tens digit out
	hundredsDigit := number / 100
	tmpNumber := number - (hundredsDigit * 100)
	tensDigit := tmpNumber / 10
	onesDigit := tmpNumber % 10
	tensWord := ""
	onesWord := ""

	// Handle single digit numbers
	ones := map[uint]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
		0: "zero",
	}

	onesWord = ones[onesDigit]
	if number < 10 {
		return onesWord, nil
	}

	// Handle teens
	if number < 20 {
		teens := map[uint]string{
			10: "ten",
			11: "eleven",
			12: "twelve",
			13: "thirteen",
			14: "fourteen",
			15: "fifteen",
			16: "sixteen",
			17: "seventeen",
			18: "eighteen",
			19: "nineteen",
		}
		return teens[number], nil
	}

	tensMap := map[uint]string{
		2: "twenty-",
		3: "thirty-",
		4: "forty-",
		5: "fifty-",
		6: "sixty-",
		7: "seventy-",
		8: "eighty-",
		9: "ninety-",
	}

	if tensDigit >= 2 {
		tensWord = tensMap[tensDigit]
	}

	if onesWord == "zero" {
		onesWord = ""
		tensWord = strings.Replace(tensWord, "-", "", 1)
	}

	if number < 100 {
		fullWord := fmt.Sprintf("%s%s", tensWord, onesWord)
		return fullWord, nil
	}

	hundredsMap := map[uint]string{
		1: "one-hundred",
		2: "two-hundred",
		3: "three-hundred",
		4: "four-hundred",
		5: "five-hundred",
		6: "six-hundred",
		7: "seven-hundred",
		8: "eight-hundred",
		9: "nine-hundred",
	}

	hundredsWord := hundredsMap[hundredsDigit]
	fullWord := fmt.Sprintf("%s %s%s", hundredsWord, tensWord, onesWord)
	fullWord = strings.TrimRight(fullWord, " ")
	return fullWord, nil
}
