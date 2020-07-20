package hamming

import "errors"

// Distance function receives two strands of DNA and calculate the hamming distance

func Distance(a, b string) (int, error) {
	counter := 0

	if len(a) != len(b) {
		return 0, errors.New("The length of the strands are different")
	} else if a == "" || b == "" {
		return 0, nil
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				counter++
			}
		}
		return counter, nil
	}
}