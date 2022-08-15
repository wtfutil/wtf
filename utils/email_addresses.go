package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// NameFromEmail takes an email address and returns the part that comes before the @ symbol
//
// Example:
//
//	NameFromEmail("test_user@example.com")
//	> "Test_user"
func NameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	name := strings.ReplaceAll(parts[0], ".", " ")

	c := cases.Title(language.English)
	return c.String(name)
}

// NamesFromEmails takes a slice of email addresses and returns a slice of the parts that
// come before the @ symbol
//
// Example:
//
//	NamesFromEmail("test_user@example.com", "other_user@example.com")
//	> []string{"Test_user", "Other_user"}
func NamesFromEmails(emails []string) []string {
	names := make([]string, len(emails))

	for i, email := range emails {
		names[i] = NameFromEmail(email)
	}

	return names
}
