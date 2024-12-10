package entropy

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"aaa", 0.0},                 // All identical characters
		{"abcd", 2.0},                // Uniform distribution
		{"www.google.com", 3.180832}, // example
	}

	for _, test := range tests {
		result := Calculate(test.input)
		if math.Abs(result-test.expected) > 0.0001 {
			t.Errorf("For input %q, expected %.6f, but got %.6f", test.input, test.expected, result)
		}
	}
}
