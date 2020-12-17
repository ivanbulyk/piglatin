package main

import (
	"fmt"
)

const (
	inputLiteral       string = "Yalantis is a great school, so are the people who work there"
	inputLiteralDigits string = "Y1l1nt3s 3s 1 gr21t sch44l, s4 1r2 th2 p24pl2 wh4 w4rk th2r2"
)

func encode(inputString string) string {
	var outputString string

	for _, symbol := range inputString {
		r := rune(symbol)

		if string(r) == "A" || string(r) == "a" {
			outputString = outputString + "1"
			continue
		} else if string(r) == "E" || string(r) == "e" {
			outputString = outputString + "2"
			continue

		} else if string(r) == "I" || string(r) == "i" {
			outputString = outputString + "3"
			continue

		} else if string(r) == "O" || string(r) == "o" {
			outputString = outputString + "4"
			continue

		} else if string(r) == "U" || string(r) == "u" {
			outputString = outputString + "5"
			continue

		}
		outputString = outputString + string(r)

	}
	return outputString
}

func decode(inputStringDecode string) string {
	var outputStringDecode string

	for _, symbol := range inputStringDecode {
		r := rune(symbol)

		if string(r) == "1" {
			outputStringDecode = outputStringDecode + "a"
			continue
		} else if string(r) == "2" {
			outputStringDecode = outputStringDecode + "e"
			continue

		} else if string(r) == "3" {
			outputStringDecode = outputStringDecode + "i"
			continue

		} else if string(r) == "4" {
			outputStringDecode = outputStringDecode + "o"
			continue

		} else if string(r) == "5" {
			outputStringDecode = outputStringDecode + "u"
			continue

		}
		outputStringDecode = outputStringDecode + string(r)

	}
	return outputStringDecode
}

func main() {
	fmt.Println("Encoding: ")
	fmt.Println(inputLiteral)
	fmt.Println(encode(inputLiteral))
	fmt.Println("Decoding: ")
	fmt.Println(decode(inputLiteralDigits))
}
