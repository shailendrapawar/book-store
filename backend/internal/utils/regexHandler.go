package utils

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsEmail(identifier string) bool {
	return emailRegex.MatchString(identifier)
}

func IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

// ==============ISBN CHECK================================
func IsISBN(s string) bool {
	cleaned := strings.ReplaceAll(s, "-", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")

	switch len(cleaned) {
	case 10:
		return isValidISBN10(cleaned)
	case 13:
		return isValidISBN13(cleaned)
	default:
		return false
	}
}

// ISBN-10 checksum validation
func isValidISBN10(s string) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
		sum += int(s[i]-'0') * (10 - i)
	}
	// last char can be 'X' = 10
	last := s[9]
	if last == 'X' || last == 'x' {
		sum += 10
	} else if unicode.IsDigit(rune(last)) {
		sum += int(last - '0')
	} else {
		return false
	}
	return sum%11 == 0
}

// ISBN-13 checksum validation
func isValidISBN13(s string) bool {
	sum := 0
	for i := 0; i < 12; i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
		if i%2 == 0 {
			sum += int(s[i] - '0')
		} else {
			sum += int(s[i]-'0') * 3
		}
	}
	check := (10 - (sum % 10)) % 10
	return check == int(s[12]-'0')
}

//=========================================================
