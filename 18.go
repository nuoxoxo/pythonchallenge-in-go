package main

import (
    "fmt"
    "os"
    "compress/gzip"
    _"strconv"
    "bytes"
    "encoding/hex"
    "strings"
    "reflect"

    "github.com/pmezard/go-difflib/difflib"
)

var buffer *bytes.Buffer

func main() {

    //  observation/ 3 slots of 32/sp in the middle
    //      len/108 ...
    //      52 / 48 uint8
    //      53 / 32 uint8
    //      54 / 32 uint8
    //      55 / 32 uint8
    //      56 / 56 uint8

    var L, R []string
    lines := strings.Split(buffer.String(), "\n")
    for _, line := range lines {
        if len(line) < 56 {
            fmt.Println("empty?/", len(line))
            //L = append(L, line)
            //continue
        } else {
            L = append(L, line[:53])
            R = append(R, line[56:])
        }
    }

    // create images
    p1, _ := os.Create("p1.png")
    defer p1.Close()

    p2, _ := os.Create("p2.png")
    defer p2.Close()

    p3, _ := os.Create("p3.png")
    defer p3.Close()

    diff, _ := difflib.GetUnifiedDiffString( difflib.UnifiedDiff {
        A:  difflib.SplitLines(strings.Join(L, "\n")),
        B:  difflib.SplitLines(strings.Join(R, "\n")),
        Context: len(L), // <--- FIXME/ bugfix: 3 is too small thats all
    })

    // print out the whole DIFF sequence
    fmt.Println(diff, "\nend/")

    for i, line := range strings.Split(diff, "\n") {
        //if i == 100 {break}
        if len(line) == 0 {
            fmt.Println("line/", i, "(null)")
            continue
        }

        bytes := []byte{}
        if !strings.HasPrefix(line, "@") {
            data := line[1:]
            bytes, _ = hexpair(data)
        }

        if strings.HasPrefix(line, "+") {
            p1.Write(bytes)
        } else if strings.HasPrefix(line, "-") {
            p2.Write(bytes)
        } else if strings.HasPrefix(line, " ") {
            //fmt.Println(i, bytes)
            p3.Write(bytes)
        } else if strings.HasPrefix(line, "@"){
            fmt.Println(i, "@ lines/", line)
        } else {
            fmt.Println(i, "default/", line)
        }
    }
}

func hexpair(hexstring string) ([]byte, error) {

    res := strings.ReplaceAll(hexstring, " ", "")
    if len(res) % 2 != 0 {
        res = "0" + res
    }
    return hex.DecodeString(res)
}

func init() {

    file, _ := os.Open("files/deltas.gz")
    defer file.Close()

    gzReader, _ := gzip.NewReader(file)
    defer gzReader.Close()

    buffer = new(bytes.Buffer)
    _, _ = buffer.ReadFrom(gzReader)

    fmt.Println(reflect.TypeOf( buffer ))
}


