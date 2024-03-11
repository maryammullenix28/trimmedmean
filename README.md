# Create a Go Package
##  How I made it & my main code

## Using trimmedmean within a Go program
This package provides a function to calculate the trimmed mean of a slice of integers or floats in Go.
### How it works
The TrimmedMean function sorts the input slice in ascending order and then trims a specified proportion of elements from the lower and upper ends. It then calculates the mean of the remaining elements and returns it.
#### Function
`func TrimmedMean(data interface{}, degreeOfTrimming ...float64) float64`
##### Parameters
- data: A slice of integers ([]int64) or floats ([]float64) for which the trimmed mean needs to be calculated. Other types will result in a zero return value.
- degreeOfTrimming: Optional. Represents the proportions of slice elements to be taken from the lower and upper ends of the sorted slice. If only one argument is provided, it is assumed to be symmetric trimming. If no argument is provided, default is set to 0.1.
##### Return Value
Returns the trimmed mean of the input slice as a float64.
