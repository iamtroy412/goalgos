package main

import (
	"fmt"
	"math"
)

func bs_list(haystack []int, needle int) bool {
    lo := 0
    hi := len(haystack)

    for (lo < hi) {
        m := int(math.Floor(float64(lo + (hi - lo) / 2)))
        v := haystack[m]

        if v == needle {
            return true
        } else if (v > needle) {
            hi = m
        } else {
            lo = m + 1
        }
    }
    return false
}

func main() {
    haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    needle1 := 42
    needle2 := 4

    fmt.Println(bs_list(haystack, needle1))
    fmt.Println(bs_list(haystack, needle2))
}
