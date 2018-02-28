// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	Acro := strings.ToUpper(string(s[0]))
	recValue := false

	for i := 1; i < len(s); i++ {
		if s[i] == 32 || s[i] == 45 {
			recValue = true
			continue
		}
		if recValue == true {
			Acro = strings.ToUpper(string(Acro)) + strings.ToUpper(string(s[i]))
			recValue = false
		}
	}
	return Acro
}
