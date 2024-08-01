package main

import (
    "fmt"
    "strconv"
    _"strings"
    "compress/gzip"
    "io"
    "os"
    "reflect"
    _"encoding/hex"
    "bytes"
    "bufio"
)

var block []uint8

func main(){

    reader := bytes.NewReader(block)
    scanner := bufio.NewScanner(reader)

    //scanner.Scan() // inspect the first line 
    //for i, thing := range scanner.Bytes() { fmt.Println(i, "/", thing, reflect.TypeOf(thing)) }

    //  observation/ 3 slots of 32/sp in the middle
    //      len/108 ...
    //      52 / 48 uint8
    //      53 / 32 uint8
    //      54 / 32 uint8
    //      55 / 32 uint8
    //      56 / 56 uint8

    L, R := [][]uint8{}, [][]uint8{}
    i := 0
    for scanner.Scan() {
        i++
        line := scanner.Bytes()
        l, r := line[:53], line[56:]
        L = append(L, l)
        R = append(R, r)
    }

    // DBG
    fmt.Println("i/", i) // total lines: 2291
    fmt.Println("dbg/", "L -", len(L), reflect.TypeOf(L), "len/0", len(L[0]))
    fmt.Println("dbg/", "L/last", L[len(L) - 1], "len/-1", len(L[len(L) - 1]))
    fmt.Println("dbg/", "R -", len(R), reflect.TypeOf(R), "len/0", len(R[0]))

    // step: diff
    //inA, inB, Both := [][]byte{}, [][]byte{}, [][]byte{}
    inA, inB, Both := []byte{}, []byte{}, []byte{}
    N := len(L)
    if N != len(R) {
        panic("different lengths/" + strconv.Itoa(N) + ":" + strconv.Itoa(len(R)))
    }
    i = 0
    for i < N {
        l, r := L[i], R[i]
        ma, mb := make(map[byte]bool), make(map[byte]bool)
        for _, val := range l {
            ma[val] = true
        }
        for _, val := range r {
            mb[val] = true
        }

        for _, val := range l {
            if mb[val] {
                Both = append(Both, val)
            } else {
                inA = append(inA, val)
            }
        }
        for _, val := range r {
            if ! ma[val] {
                inB = append(inB, val)
            }
        }
        i++
    }
    fmt.Println("len only a/", len(inA))
    fmt.Println("len only b/", len(inB))
    fmt.Println("len Both/", len(Both))

    img1, _ := os.Create("out1")
    defer img1.Close()
    img1.Write(inA)
}


func init(){
    file, _ := os.Open("files/deltas.gz")
    defer file.Close()
    reader, _ := gzip.NewReader(file)
    defer reader.Close()
    data, _ := io.ReadAll(reader)
    block = data
}


