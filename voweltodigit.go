package main

import (
	"fmt"
)

const (
	inputEncode string = "Yalantis is a great school, so are the people who work there"
	inputDecode string = "Y1l1nt3s 3s 1 gr21t sch44l, s4 1r2 th2 p24pl2 wh4 w4rk th2r2"
)

var (
	match = map[string]string{
		"a": "1",
		"e": "2",
		"i": "3",
		"o": "4",
		"u": "5",
		"A": "1",
		"E": "2",
		"I": "3",
		"O": "4",
		"U": "5",
	}

	matchBack = map[string]string{
		"1": "a",
		"2": "e",
		"3": "i",
		"4": "o",
		"5": "u",
	}
)

func encode(inputString string) string {
	var outputString string

outer:

	for _, symbol := range inputString {
		r := rune(symbol)
		for key, value := range match {
			if string(r) == key {
				outputString = outputString + value
				continue outer
			}
		}

		outputString = outputString + string(r)
	}
	return outputString
}

func decode(inputStringDecode string) string {
	var outputStringDecode string

outer:

	for _, symbol := range inputStringDecode {
		r := rune(symbol)
		for key, value := range matchBack {
			if string(r) == key {
				outputStringDecode = outputStringDecode + value
				continue outer
			}
		}

		outputStringDecode = outputStringDecode + string(r)
	}
	return outputStringDecode
}

func main() {
	fmt.Println("Encoding: ")
	fmt.Println(inputEncode)
	fmt.Println(encode(inputEncode))
	fmt.Println("Decoding: ")
	fmt.Println(decode(inputDecode))
}
