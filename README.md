# Create a Go Package
## Trimmed Mean Package
### Using trimmedmean within a Go program
This package provides a function to calculate the trimmed mean of a slice of integers or floats in Go.
#### Example
```
package main

import (
    "fmt"
    "github.com/maryammullenix28/trimmedmean"
)

func main() {
    data := []float64{10, 20, 30, 40, 50}
    trimmedMean := trimmedmean.TrimmedMean(data, 0.2) // Calculate trimmed mean with symmetric 20% trimming
    fmt.Println("Trimmed Mean:", trimmedMean)
}
```

### How it works
The TrimmedMean function sorts the input slice in ascending order and then trims a specified proportion of elements from the lower and upper ends. It then calculates the mean of the remaining elements and returns it.
#### Function
`func TrimmedMean(data interface{}, degreeOfTrimming ...float64) float64`
##### Parameters
- data: A slice of integers ([]int64) or floats ([]float64) for which the trimmed mean needs to be calculated. Other types will result in a zero return value.
- degreeOfTrimming: Optional. Represents the proportions of slice elements to be taken from the lower and upper ends of the sorted slice. If only one argument is provided, it is assumed to be symmetric trimming. If no argument is provided, default is set to 0.1.
##### Return Value
Returns the trimmed mean of the input slice as a float64.
## Implementation of `main`
The code utilizes the following packages:
- "fmt"
- "math/rand": To generate random ints and floats for test data
- "strings"
- "trimmedmean": Assignement package
- ["github.com/senseyeio/roger"](https://pkg.go.dev/github.com/senseyeio/roger): Roger is a Go RServe client, allowing the capabilities of R to be used from Go applications.
## Running the `main` program
Before running the Go program, open and run RserveClient.R. This will start the R client used in `main` to run the R code on the same set of ints and floats.
