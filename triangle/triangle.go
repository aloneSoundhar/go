// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT = 0
	Equ = 1
	Iso = 2
	Sca = 3
)

// ShareWith should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind
	if (a <= 0 || b <= 0 || c <= 0) || !(a+b >= c && b+c >= a && c+a >= b) {
		k = NaT
	} else {
		if a == b && b == c {
			k = Equ
		} else if a == b || b == c || c == a {
			k = Iso
		} else {
			k = Sca
		}
	}
	return k
}
