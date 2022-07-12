package main

import (
	"fmt"
	"testing"
)

func TestGetPrimeFactors(t *testing.T) {
	if fmt.Sprintf("%v", getPrimeFactors(17)) != `[17]` {
		t.Error(17)
	}
	if fmt.Sprintf("%v", getPrimeFactors(12)) != `[2 2 3]` {
		t.Error(12)
	}
	if fmt.Sprintf("%v", getPrimeFactors(360)) != `[2 2 2 3 3 5]` {
		t.Error(360)
	}
	if fmt.Sprintf("%v", getPrimeFactors(1)) != `[]` {
		t.Error(1)
	}
}
