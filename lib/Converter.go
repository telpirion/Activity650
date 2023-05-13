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
		1: "one",
	}

	return intsToWords[number], nil
}
