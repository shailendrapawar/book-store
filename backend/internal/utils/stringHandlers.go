package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NormalizeISBN(isbn string) string {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbn = strings.ReplaceAll(isbn, " ", "")
	return isbn
}

func CreateUUID() string {
	return uuid.New().String()
}
