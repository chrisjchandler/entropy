package entropy

import "testing"
import "math"

func TestCalculate(t *testing.T) {
    input := "hello"
    expected := 1.921928 // Correct entropy for "hello"
    result := Calculate(input)
    if math.Abs(result-expected) > 1e-6 {
        t.Errorf("Expected %f, got %f", expected, result)
    }
}


