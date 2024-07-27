package main

import (
    "fmt"
)

func main(){

    var n int64 = 1
    i := 0
    for i < 38 {
        n *= 2
        i++
    }
    fmt.Println(n)
}
