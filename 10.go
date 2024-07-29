package main

import (
    "fmt"
    "strconv"
)

func main(){
    sequence := "1"
    i := 0
    for i < 30 {
        temp := ""
        fast, slow := 0, 0
        for slow < len(sequence) {
            char := string(sequence[slow])
            for fast < len(sequence) && sequence[fast] == sequence[slow] {
                fast++
            }
            diff := fast - slow
            slow = fast
            times := strconv.Itoa(diff)
            temp += times + char
        }
        N := len(sequence)
        if N <= 42 {
            fmt.Println(i, "/", sequence)
        } else {
            fmt.Println(i, "/ len/", N)
        }
        sequence = temp
        i++
    }
    fmt.Println(/*"end/", sequence,*/ "len/", len(sequence))
}


