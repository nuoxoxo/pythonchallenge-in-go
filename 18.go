package main

import (
    "fmt"
    "strconv"
    "strings"
    "compress/gzip"
    "io"
    "os"
    "reflect"
    "encoding/hex"
    _"bytes"
    _"image/png"
)

var lines []string

func main(){

    fmt.Println("res/")
    discrepancy := 1
    L, R := []string{}, []string{}
    for i, line := range lines {
        if len(line) < 64 {
            fmt.Println("empty?/", strconv.Quote(line), "len", len(line))
        } else {
            l, r := line[:53], line[56:]
            fmt.Println(i, line, strconv.Itoa(len(l)) + ":" + strconv.Itoa(len(r)))
            if len(l) != len(r) {
                fmt.Println("error?/\t", discrepancy)
                discrepancy++
            }
            L, R = append(L, l), append(R, r)
        }
    }

    // DBG
    test := 100
    testL, testR := L[test], R[test]
    fmt.Println("dbg/", len(L), reflect.TypeOf(L), testL, len(testL), testL[len(testL) - 1])
    fmt.Println("dbg/", len(R), reflect.TypeOf(R), testR, len(testR), testR[len(testR) - 1])

    // []string ---> []bytes
    bytesL, bytesR := make([][]byte, len(L)), make([][]byte, len(R))
    for i, line := range L {
        pairs := strings.Split(line, " ")
        fmt.Println(len(pairs))
        temp := make([]byte, len(pairs))
        for j, pair := range pairs {
            val, _ := hex.DecodeString(pair)
            if len(val) > 0 {
                temp[j] = val[0]
            } else {
                temp[j] = 0x00
            }
        }
        bytesL[i] = temp
    }
    for i, line := range R {
        pairs := strings.Split(line, " ")
        if len(pairs) == 0 {continue}
        temp := make([]byte, len(pairs))
        for j, pair := range pairs {
            val, _ := hex.DecodeString(pair)
            if len(val) > 0 {
                temp[j] = val[0]
            } else {
                temp[j] = 0x00
            }
        }
        bytesR[i] = temp
    }
    fmt.Println(bytesL[0], bytesL[42])
    fmt.Println(bytesR[0], bytesR[42])


    // diff
    onlya, onlyb, both := [][]byte{}, [][]byte{}, [][]byte{}
    N := len(bytesL)
    if N != len(bytesR) {
        panic("different lengths/" + strconv.Itoa(N) + ":" + strconv.Itoa(len(bytesR)))
    }
    i := 0
    for i < N {
        l, r := bytesL[i], bytesR[i]
        ma, mb := make(map[byte]bool), make(map[byte]bool)
        for _, val := range l {
            ma[val] = true
        }
        for _, val := range r {
            mb[val] = true
        }

        oa, ob, bo := []byte{}, []byte{}, []byte{}
        for _, val := range l {
            if mb[val] {
                bo = append(bo, val)
            } else {
                oa = append(oa, val)
            }
        }
        for _, val := range r {
            if ma[val] { // in fact it's been done in above loop
                bo = append(bo, val)
            } else {
                ob = append(ob, val)
            }
        }
        onlya = append(onlya, oa)
        onlyb = append(onlyb, ob)
        both = append(both, bo)
        i++
    }
    fmt.Println("len only a/", len(onlya))
    fmt.Println("len only b/", len(onlyb))
    fmt.Println("len both/", len(both))

    /*
    aData := []byte{}
    for _, data := range onlya {
        aData = append(aData, data...)
    }
    fmt.Println(aData)
    reader := bytes.NewReader(aData)
    img1, err := png.Decode(reader)
    if err != nil {
        fmt.Println("err/", err)
    }
    img1out, _ := os.Create("img1.png")
    defer img1out.Close()
    png.Encode(img1out, img1)
    */
    img1, _ := os.Create("1.png")
    defer img1.Close()
    for i, line := range onlya[:2218] {
        _, err := img1.Write(line)
        fmt.Println(i, line, len(line))
        if err != nil {fmt.Println(err)}
    }
}


func init(){
    file, _ := os.Open("files/deltas.gz")
    defer file.Close()
    reader, _ := gzip.NewReader(file)
    defer reader.Close()
    data, _ := io.ReadAll(reader)
    lines = strings.Split(string(data), "\n")
}


