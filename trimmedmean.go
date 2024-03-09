package trimmedmean

func Average(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
