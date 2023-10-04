package main

import "fmt"

func foo(n int) int {
    if n == 1 {
        fmt.Println(n)
        return 1
    }

    out := n + foo(n - 1)
    fmt.Println(n)
    return out
}

func main() {
    fmt.Println(foo(5)) 
}
