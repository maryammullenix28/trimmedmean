package trimmedmean

import "sort"

// TrimmedMean calculates the trimmed mean of a slice of integers or floats.
// The degreeOfTrimming arguments represent the proportions of slice elements
// to be taken from the lower and upper ends of the slice after sorting.
// If only one degreeOfTrimming argument is provided, it is assumed that the
// trimming is symmetric.
func TrimmedMean(data []float64, degreeOfTrimming ...float64) float64 {
	if len(data) == 0 {
		return 0
	}

	// Sort the data
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	// Calculate the number of elements to trim from each end
	var trimLower, trimUpper int
	if len(degreeOfTrimming) == 0 {
		// Default symmetric trimming
		trim := int(float64(len(sortedData)) * 0.1) // Default trimming proportion
		trimLower, trimUpper = trim, trim
	} else if len(degreeOfTrimming) == 1 {
		// Symmetric trimming
		trim := int(float64(len(sortedData)) * degreeOfTrimming[0])
		trimLower, trimUpper = trim, trim
	} else {
		// Asymmetric trimming
		trimLower = int(float64(len(sortedData)) * degreeOfTrimming[0])
		trimUpper = int(float64(len(sortedData)) * degreeOfTrimming[1])
	}

	// Calculate the sum of trimmed elements
	sum := 0.0
	for i := trimLower; i < len(sortedData)-trimUpper; i++ {
		sum += sortedData[i]
	}

	// Calculate and return the trimmed mean
	return sum / float64(len(sortedData)-trimLower-trimUpper)
}
