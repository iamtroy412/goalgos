package main

import "fmt"

func bubble_sort(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr) - 1 - i; j++ {
            if arr[j] > arr[j+1] {
                // t := arr[j]
                // arr[j] = arr[j+1]
                // arr[j+1] = t
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}


func main() {
    arr := []int{1, 3, 4, 7, 2}
    arr2 := []int{1, 3, 4, 7, 2, 42, 69, 8, 11, 42, 100, 1, 99}
    bubble_sort(arr)
    bubble_sort(arr2)
    fmt.Println(arr)
    fmt.Println(arr2)
}
