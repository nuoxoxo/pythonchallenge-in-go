package main

import (
    "fmt"
    "strconv"
    _"strings"
    "compress/gzip"
    "io"
    "os"
    "reflect"
    "encoding/hex"
    "bytes"
    "bufio"
    _"image"
    _"image/color"
    _"image/png"
    "io/ioutil"
)

var block []uint8

func main(){

    reader := bytes.NewReader(block)
    scanner := bufio.NewScanner(reader)

    // scanner.Scan() // inspect the first line 
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
        ll, rr := line[:53], line[56:]

        // added
        //ll = bytes.TrimSpace(ll)
        //rr = bytes.TrimSpace(rr)

        l, _ := hex.DecodeString(string(ll))
        r, _ := hex.DecodeString(string(rr))

        //fmt.Println(len(l), ":", len(r))
        L = append(L, l)
        R = append(R, r)
    }

    // DBG
    fmt.Println("i/", i) // total lines: 2291
    fmt.Println("dbg/", "L -", len(L), reflect.TypeOf(L), "len/0", len(L[0]))
    fmt.Println("dbg/", "L/last", L[len(L) - 1], "len/-1", len(L[len(L) - 1]))
    fmt.Println("dbg/", "R -", len(R), reflect.TypeOf(R), "len/0", len(R[0]))

    // step: diff
    //A, B, AB := [][]byte{}, [][]byte{}, [][]byte{}
    A, B, AB := []byte{}, []byte{}, []byte{}
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
                AB = append(AB, val)
            } else {
                A = append(A, val)
            }
        }
        for _, val := range r {
            if ! ma[val] {
                B = append(B, val)
            }
        }
        i++
    }
    fmt.Println("len only a/", len(A))
    fmt.Println("len only b/", len(B))
    fmt.Println("len AB/", len(AB))

    out1, _ := os.Create("out1.png")
    defer out1.Close()
    _ = ioutil.WriteFile("out1.png", A, 0644)

}


func init(){
    file, _ := os.Open("files/deltas.gz")
    defer file.Close()
    reader, _ := gzip.NewReader(file)
    defer reader.Close()
    data, _ := io.ReadAll(reader)
    block = data
}


