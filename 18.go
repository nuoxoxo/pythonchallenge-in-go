package main

import (
    _"strconv"
    "bytes"
    "compress/gzip"
    "encoding/hex"
    "fmt"
    _"io"
    "os"
    "strings"

    "github.com/pmezard/go-difflib/difflib"
)

func main() {

    file, _ := os.Open("files/deltas.gz")
    defer file.Close()

    gzReader, _ := gzip.NewReader(file)
    defer gzReader.Close()

    var L, R []string
    buffer := new(bytes.Buffer)
    _, _ = buffer.ReadFrom(gzReader)

    lines := strings.Split(buffer.String(), "\n")
    for _, line := range lines {
        if len(line) < 56 {
            fmt.Println("56 ?/", len(line))
        } else {
            L = append(L, line[:53])
            R = append(R, line[56:])
        }
    }

    diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
        A: difflib.SplitLines(strings.Join(L, "\n")),
        B: difflib.SplitLines(strings.Join(R, "\n")),
        Context: 3,
    })

    p1, _ := os.Create("p1.png")
    defer p1.Close()

    p2, _ := os.Create("p2.png")
    defer p2.Close()

    p3, _ := os.Create("p3.png")
    defer p3.Close()

    //fmt.Println(diff)

    for i, line := range strings.Split(diff, "\n") {
        if len(line) == 0 {
            continue
        }

        var bs []byte
        if !strings.HasPrefix(line, "@") {//&& (strings.HasPrefix(line, "+") || strings.HasPrefix(line, "-")) {
            data := line[1:]
            fmt.Println(i, line)
            bs, _ = hexstring2bytes(data)
            //data := strings.TrimSpace(line[1:])
            //bs, _ = hexstring2bytes(data)
        }
        //if len(bs) < 1 { continue } 
        //switch {
        if strings.HasPrefix(line, "+") {
            p2.Write(bs)
            //fmt.Println("+/", line)
        } else if strings.HasPrefix(line, "-") {
            p3.Write(bs)
            //fmt.Println("-/", strconv.Quote(line))
        } else if strings.HasPrefix(line, " "){
            p1.Write(bs)
            fmt.Println(i, "prefix!=@/", len(bs), line)
        } 
        /*else {
            fmt.Println(i, "default/", line)
        }*/
        //default:
        //fmt.Println(i, "default/", line)
        }
    }

func hexstring2bytes(hexstring string) ([]byte, error) {
    hexstring = strings.ReplaceAll(hexstring, " ", "")
    if len(hexstring) % 2 != 0 {
        hexstring = "0" + hexstring
    }
    return hex.DecodeString(hexstring)
}

