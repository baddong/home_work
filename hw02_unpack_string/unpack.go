package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString returns if input string contains invalid characters
var ErrInvalidString = errors.New("invalid string")

// Unpack formats the string according to the given format: "a4bc2d5e" => "aaaabccddddde"
func Unpack(input string) (string, error) {
	var unpackedString strings.Builder
	var repeatsCount int
	var repeatedCh string
	var prevIsDigit bool

	input += "\n" // to successfully iterate over last char in "for" loop

	if unicode.IsDigit(rune(input[0])) {
		return "", ErrInvalidString
	}

	// in this loop I use WriteString after I looked on char in next iteration
	for _, ch := range input {
		repeatsCount, _ = strconv.Atoi(string(ch))
		switch {
		case unicode.IsDigit(ch):
			if prevIsDigit {
				return "", ErrInvalidString
			}
			prevIsDigit = true
		case prevIsDigit:
			repeatedCh = string(ch)
			prevIsDigit = false
			continue
		case !prevIsDigit:
			prevIsDigit = false
			repeatsCount = 1
		}
		unpackedString.WriteString(strings.Repeat(repeatedCh, repeatsCount))
		repeatedCh = string(ch)
	}

	result := unpackedString.String()
	return result, nil
}
