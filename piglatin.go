package main

import (
	"fmt"
	"strings"
	"unicode"
)

func translate(inputWords string) string {
	var piglatinWords []string

	englishWords := strings.Split(inputWords, " ")

	fmt.Printf("Value of englishWords:   %v\n", englishWords)
	fmt.Println()

	for _, word := range englishWords {
		if len(word) < 2 && unicode.IsPunct([]rune(word[0:1])[0]) {
			piglatinWords = append(piglatinWords, word)
			continue
		} else if len(word) < 3 {
			word = word + string(rune(0)) + string(rune(0)) + string(rune(0))
		}
		firstLetter := word[0:1]
		secondLetter := word[1:2]
		thirdLetter := word[2:3]
		lastLetter := word[len(word)-1:]
		r := []rune(firstLetter)
		rLast := []rune(lastLetter)

		switch {
		case unicode.IsPunct(r[0]): //case when first sign IS NOT a letter
			changeSignFirst := word[0:1]
			word = word[1:]
			switch {
			case unicode.IsPunct(rLast[0]): //case when first sign IS NOT a letter ++ case when AND LAST sign IS NOT a letter TOO
				changeSignLast := word[len(word)-1:]
				word = word[:len(word)-1]
				//can be another first letter here

				//maybe
				switch firstLetter {
				case "a", "e", "i", "o", "u": //case when first letter is vowel
					piglatinWords = append(piglatinWords, changeSignFirst+word+"ay"+changeSignLast)
				default: //case when first letter is consonant
					switch secondLetter {
					case "a", "e", "i", "o", "u":
						piglatinWords = append(piglatinWords, changeSignFirst+word[1:]+word[0:1]+"ay"+changeSignLast)
					default: //case when second letter is consonant
						switch thirdLetter {
						case "a", "e", "i", "o", "u":
							piglatinWords = append(piglatinWords, changeSignFirst+word[2:]+word[0:2]+"ay"+changeSignLast)
						default: //case when third letter is consonant
							piglatinWords = append(piglatinWords, changeSignFirst+word[3:]+word[0:3]+"ay"+changeSignLast)
						}

					}

				}
			default:
				switch firstLetter {
				case "a", "e", "i", "o", "u": //case when first letter is vowel
					piglatinWords = append(piglatinWords, changeSignFirst+word+"ay")
				default: //case when first letter is consonant
					switch secondLetter {
					case "a", "e", "i", "o", "u":
						piglatinWords = append(piglatinWords, changeSignFirst+word[1:]+word[0:1]+"ay")
					default: //case when second letter is consonant
						switch thirdLetter {
						case "a", "e", "i", "o", "u":
							piglatinWords = append(piglatinWords, changeSignFirst+word[2:]+word[0:2]+"ay")
						default: //case when third letter is consonant
							piglatinWords = append(piglatinWords, changeSignFirst+word[3:]+word[0:3]+"ay")
						}

					}

				}
			}
		case unicode.IsPunct(rLast[0]): //case when LAST sign IS NOT a letter
			changeSignLast := word[len(word)-1:]
			word = word[:len(word)-1]
			switch firstLetter {
			case "a", "e", "i", "o", "u": //case when first letter is vowel
				piglatinWords = append(piglatinWords, word+"ay"+changeSignLast)
			default: //case when first letter is consonant
				switch secondLetter {
				case "a", "e", "i", "o", "u":
					piglatinWords = append(piglatinWords, word[1:]+word[0:1]+"ay"+changeSignLast)
				default: //case when second letter is consonant
					switch thirdLetter {
					case "a", "e", "i", "o", "u":
						piglatinWords = append(piglatinWords, word[2:]+word[0:2]+"ay"+changeSignLast)
					default: //case when third letter is consonant
						piglatinWords = append(piglatinWords, word[3:]+word[0:3]+"ay"+changeSignLast)
					}

				}

			}
		default:
			switch firstLetter {
			case "a", "e", "i", "o", "u": //case when first letter is vowel
				piglatinWords = append(piglatinWords, word+"ay")
			default: //case when first letter is consonant
				switch secondLetter {
				case "a", "e", "i", "o", "u":
					piglatinWords = append(piglatinWords, word[1:]+word[0:1]+"ay")
				default: //case when second letter is consonant
					switch thirdLetter {
					case "a", "e", "i", "o", "u":
						piglatinWords = append(piglatinWords, word[2:]+word[0:2]+"ay")
					default: //case when third letter is consonant
						piglatinWords = append(piglatinWords, word[3:]+word[0:3]+"ay")
					}

				}

			}
		}

	}

	return strings.Join(piglatinWords, " ")
}

func main() {
	fmt.Println(translate("Yalantis , is a great school, so are the people ,who work there"))
}
