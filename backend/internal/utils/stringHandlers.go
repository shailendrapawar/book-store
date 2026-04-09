package utils

import (
	"strings"

	"github.com/aarondl/opt/null"
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

func ExtractNullString(ns null.Val[string]) string {
	if val, ok := ns.Get(); ok {
		return val
	}
	return ""
}
