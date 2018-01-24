package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Input strings must be the same length.")
	}

	var hamming int = 0
	for ii:=0; ii < len(a); ii++ {
		if (b[ii] != a[ii]) {
			hamming ++
		}
	}
	return hamming, nil

}
