package hamming

import (
	"errors"
)

// Distance function receives two strands of DNA and calculate the hamming distance
// expects two strings (a,b) and returns an int indicating the hamming distance and an
// error if something not gone as expected like receiving two strands of different length.
func Distance(a, b string) (int, error) {
	counter := 0
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("The length of the strands are different")
	}
	for i := range ar {
		if ar[i] != br[i] {
			counter++
		}
	}
	return counter, nil
}
