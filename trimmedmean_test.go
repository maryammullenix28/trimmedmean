package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	// Empty slice
	data := []float64{}
	expected := 0.0
	result := TrimmedMean(data)
	if result != expected {
		t.Errorf("Test case 1 failed: Expected %f, got %f", expected, result)
	}

	// Symmetric trimming with 10% from each end
	data = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected = 5.5
	result = TrimmedMean(data, 0.1)
	if result != expected {
		t.Errorf("Test case 2 failed: Expected %f, got %f", expected, result)
	}

	// Asymmetric trimming with 20% from the lower end and 10% from the upper end
	data = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected = 6
	result = TrimmedMean(data, 0.2, 0.1)
	if result != expected {
		t.Errorf("Test case 3 failed: Expected %f, got %f", expected, result)
	}

}
