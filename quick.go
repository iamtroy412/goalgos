package main

import "fmt"

func qs(arr *[]int, low, hi int) {
    if low >= hi {
        return
    }

    pivotIdx := parition(arr, low, hi)

    qs(arr, low, pivotIdx - 1)
    qs(arr, pivotIdx + 1, hi)
}

func parition(arr *[]int, low, hi int) int {
    pivot := (*arr)[hi]

    idx := low - 1

    for i := low; i < hi; i++ {
        if (*arr)[i] <= pivot {
            idx++
            tmp := (*arr)[i]
            (*arr)[i] = (*arr)[idx]
            (*arr)[idx] = tmp
        }
    }

    idx++
    (*arr)[hi] = (*arr)[idx]
    (*arr)[idx] = pivot
    return idx
}

func quick_sort(arr *[]int) {
    qs(arr, 0, len(*arr) - 1)
}

func main() {
    array := []int{9, 3, 7, 4, 69, 420, 42}

    fmt.Println("Before:", array)
    quick_sort(&array)
    fmt.Println("After:", array)
}
