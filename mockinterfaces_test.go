package mockinterfaces_test

import (
	"testing"

	fib "github.com/brentgroves/go_test_tips"
)

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := fib.Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
