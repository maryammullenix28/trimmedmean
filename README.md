# Create a Go Package
##  How I made it
## Using trimmedmean within a Go program
#### Function
`func TrimmedMean(data interface{}, degreeOfTrimming ...float64) float64`
##### Parameters
- data: A slice of integers ([]int64) or floats ([]float64) for which the trimmed mean needs to be calculated.
- degreeOfTrimming: Optional. Represents the proportions of slice elements to be taken from the lower and upper ends of the sorted slice. If only one argument is provided, it is assumed to be symmetric trimming.
