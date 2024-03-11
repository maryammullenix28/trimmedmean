package main

import (
	"fmt"
	"math/rand"
	"strings"
	"trimmedmean"

	"github.com/senseyeio/roger"
)

func main() {
	// Seed the random number generator
	rand.Seed(100)

	// Generate sample data
	intData := generateRandomIntData(100)
	floatData := generateRandomFloatData(100)

	// Compute trimmed mean using trimmedmean package
	goIntTrimmedMean := trimmedmean.TrimmedMean(intData, 0.05)
	goFloatTrimmedMean := trimmedmean.TrimmedMean(floatData, 0.05)

	// R CODE FROM HERE
	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		fmt.Println("Failed to connect to R:", err)
		return
	}

	// Convert arrays to comma-separated strings
	intDataStr := convertToRVector(intData)
	floatDataStr := convertToRVector(floatData)

	// Compute trimmed means using R
	intTrimmedMean, err := rClient.Eval(fmt.Sprintf("mean(c(%s), trim=0.05)", intDataStr))
	if err != nil {
		fmt.Println("Error computing trimmed mean for integer data:", err)
		return
	}
	floatTrimmedMean, err := rClient.Eval(fmt.Sprintf("mean(c(%s), trim=0.05)", floatDataStr))
	if err != nil {
		fmt.Println("Error computing trimmed mean for float data:", err)
		return
	}

	// Print results
	fmt.Println("Integer Data:")
	fmt.Println("Go Trimmed Mean:", goIntTrimmedMean)
	fmt.Println("R Trimmed Mean:", intTrimmedMean)
	fmt.Println("Float Data:")
	fmt.Println("Go Trimmed Mean:", goFloatTrimmedMean)
	fmt.Println("R Trimmed Mean:", floatTrimmedMean)
}

func generateRandomIntData(n int) []int64 {
	data := make([]int64, n)
	for i := 0; i < n; i++ {
		data[i] = int64(rand.Intn(100)) // Generate random integers between 0 and 99
	}
	return data
}

func generateRandomFloatData(n int) []float64 {
	data := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Float64() * 100 // Generate random floats between 0 and 100
	}
	return data
}

func convertToRVector(data interface{}) string {
	switch data := data.(type) {
	case []int64:
		var sb strings.Builder
		for i, val := range data {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%d", val))
		}
		return sb.String()
	case []float64:
		var sb strings.Builder
		for i, val := range data {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%f", val))
		}
		return sb.String()
	default:
		return ""
	}
}
