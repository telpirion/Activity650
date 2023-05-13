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

func IntegerToWordedString(number uint) (string, error) {
	intsToWords := map[uint]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		0:  "zero",
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

	return intsToWords[number], nil
}
