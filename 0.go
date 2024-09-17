package main

import (
    "fmt"
)

func main(){
    var n uint64 = 1
    for i := 0; i < 38; i++ {
        n *= 2
        fmt.Println(i, "\b/\t", n)
    }
    fmt.Println("res/\t", n)
}
