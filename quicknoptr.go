package main

import "fmt"

func qs2(arr []int, low, hi int) []int {
    if low >= hi {
        return arr
    }

    arr, pivotIdx := partition2(arr, low, hi)
    arr = qs2(arr, low, pivotIdx - 1)
    arr = qs2(arr, pivotIdx + 1, hi)
    return arr
}

func partition2(arr []int, low, hi int) ([]int, int) {
    pivot := arr[hi]

    idx := low - 1

    for i := low; i < hi; i++ {
        if arr[i] <= pivot {
            idx++
            tmp := arr[i]
            arr[i] = arr[idx]
            arr[idx] = tmp
        }
    }

    idx++
    arr[hi] = arr[idx]
    arr[idx] = pivot
    return arr, idx
}


func quick_sort2(arr []int) []int {
    return qs2(arr, 0, len(arr) - 1)
}

func main() {
    array2 := []int{9, 3, 7, 4, 69, 420, 42, 99}

    fmt.Println("Before:", array2)
    array2 = quick_sort2(array2)
    fmt.Println("After:", array2)
}
