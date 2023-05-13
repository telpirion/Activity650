/**
 * Eric Schmidt
 * CSS 508 Spring 2023
 * 2023-05-13
 *
 * A simple console app interface for the conversion code.
 * The console app takes a single integer parameter and outputs
 * the converted worded string value on the console.
 * This is called Program.cs in the README.
 *
 * Usage:
 * go run . -n=[NUM_TO_CONVERT]
 */
package main

import (
	"activity650/lib"
	"flag"
	"fmt"
)

func main() {
	numPtr := flag.Int("n", 0, "the number to convert to English")
	flag.Parse()
	numConvert := *numPtr

	fmt.Printf("Convert the integer %d to English:\n", numConvert)
	englishOutput, err := lib.IntegerToWordedString(uint(numConvert))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("English output for integer is: %s\n", englishOutput)
}
