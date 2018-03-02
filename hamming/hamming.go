package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	// Distance program
	count := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count, nil
	}
	return -1, errors.New("Err")
}
