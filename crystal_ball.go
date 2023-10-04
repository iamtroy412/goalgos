package main

import (
	"fmt"
	"math"
)

func two_crystal_balls(breaks []bool) int {
    jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

    i := jumpAmount
    for ; i < len(breaks); i += jumpAmount {
        if breaks[i] {
            break
        }
    }

    i -= jumpAmount;

    for j := i; j < len(breaks); j++ {
        if breaks[j] {
            return j
        }
    }

    return -1

}

func main() {
    breaks := []bool{false, false, false, true, true, true, true, true}
    breaks2 := []bool{false, false, false, false, false, false, true, true}

    fmt.Println(two_crystal_balls(breaks))
    fmt.Println(two_crystal_balls(breaks2))
}
