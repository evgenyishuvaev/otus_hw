package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString      = errors.New("invalid string")
	ErrFirstSymbolIsDigit = fmt.Errorf("string begins from digit: %w", ErrInvalidString)
	ErrNumberWasFound     = fmt.Errorf("number was found: %w", ErrInvalidString)
)

func Unpack(inputString string) (string, error) {
	var (
		resultString strings.Builder
		prevSym      rune
	)

	if inputString == "" {
		return "", nil
	}
	firstSym := []rune(inputString)[0]
	if unicode.IsDigit(firstSym) {
		return "", ErrFirstSymbolIsDigit
	}
	for _, curSym := range inputString {
		if unicode.IsDigit(prevSym) && unicode.IsDigit(curSym) {
			return "", ErrNumberWasFound
		}
		if unicode.IsLetter(prevSym) && unicode.IsLetter(curSym) {
			chunk := prevSym
			resultString.WriteRune(chunk)
		}
		if unicode.IsLetter(prevSym) && unicode.IsDigit(curSym) {
			multiplier, err := strconv.Atoi(string(curSym))
			if err != nil {
				err = fmt.Errorf("can't parse string into int: %w", err)
				return "", err
			}
			chunk := strings.Repeat(string(prevSym), multiplier)
			resultString.WriteString(chunk)
		}
		prevSym = curSym
	}
	if !unicode.IsDigit(prevSym) {
		resultString.WriteRune(prevSym)
	}
	return resultString.String(), nil
}
