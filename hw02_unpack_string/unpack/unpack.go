package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	// "strconv"
	"strings"
	// "unicode"
)

var ErrInvalidString = errors.New("invalid string")
var ErrFirstSymbolIsDigit = fmt.Errorf("string begins from digit: %w", ErrInvalidString)
var ErrNumberWasFound = fmt.Errorf("number was found: %w", ErrInvalidString)

func Unpack(inputString string) (string, error) {
	// var stringBuilder strings.Builder = strings.Builder{}
	var resultString strings.Builder
	var prevSym rune
	var chunk string
	if inputString == "" {
		return "", nil
	}

	for i, curSym := range inputString {
		if i == 0 && unicode.IsDigit(curSym) {
			return "", ErrFirstSymbolIsDigit
		}

		if unicode.IsDigit(prevSym) && unicode.IsDigit(curSym) {
			return "", ErrNumberWasFound
		}

		if unicode.IsLetter(prevSym) && unicode.IsLetter(curSym) {
			chunk = string(prevSym)
			// fmt.Println("оба строки", chunk, prevSym, curSym)
			resultString.WriteString(chunk)
		}

		if unicode.IsLetter(prevSym) && unicode.IsDigit(curSym) {
			multiplier, err := strconv.Atoi(string(curSym))
			// fmt.Println(multiplier)
			if err != nil {
				err = fmt.Errorf("can't parse string into int: %w", err)
				return "", err
			}
			chunk = strings.Repeat(string(prevSym), multiplier)
			resultString.WriteString(chunk)
			// fmt.Println("есть множитель", chunk, prevSym, curSym)
		}
		// fmt.Println(chunk)
		prevSym = curSym

	}
	resultString.WriteString(string(prevSym))
	return resultString.String(), nil
}
