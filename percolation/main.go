package main

import (
    "fmt"

    "github.com/iamtroy412/goalgos/percolation/percolationstats"
)

func main () {
    stats := percolationstats.New(10, 30)

    fmt.Printf("mean = %v\n", stats.Mean())
    fmt.Printf("stddev = %v\n", stats.StdDev())
    fmt.Printf("95 confidence high %v\n", stats.ConfidenceHigh())
    fmt.Printf("95 confidence high %v\n", stats.ConfidenceLow())
}
