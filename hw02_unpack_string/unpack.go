package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString returns if defaultString contains invalid characters
var ErrInvalidString = errors.New("invalid string")

// Unpack formats the string according to the given format: "a4bc2d5e" => "aaaabccddddde"
func Unpack(defaultString string) (string, error) {
	var unpackedString strings.Builder
	var numberOfRepeats int
	repeatedChar := ""
	var previousCharIsDigit bool = false

	defaultString += "\n" // to successfully iterate over last char in "for" loop

	// in this loop I use WriteString after I looked on char in next iteration
	for _, defaultChar := range defaultString {
		numberOfRepeats, _ = strconv.Atoi(string(defaultChar))
		switch {
		case unicode.IsDigit(defaultChar):
			if previousCharIsDigit {
				return "", ErrInvalidString
			}
			previousCharIsDigit = true
		case previousCharIsDigit:
			repeatedChar = string(defaultChar)
			previousCharIsDigit = false
			continue
		case !previousCharIsDigit:
			previousCharIsDigit = false
			numberOfRepeats = 1
		}
		unpackedString.WriteString(strings.Repeat(repeatedChar, numberOfRepeats))
		repeatedChar = string(defaultChar)
	}

	result := unpackedString.String()
	return result, nil
}
