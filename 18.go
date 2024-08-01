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
    // Open the gzip file
    file, err := os.Open("files/deltas.gz")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a gzip reader
gzReader, err := gzip.NewReader(file)
if err != nil {
fmt.Println("Error creating gzip reader:", err)
return
}
defer gzReader.Close()

var d1, d2 []string
buf := new(bytes.Buffer)
if _, err = buf.ReadFrom(gzReader); err != nil {
fmt.Println("Error reading gzip data:", err)
return
}
lines := strings.Split(buf.String(), "\n")

for _, line := range lines {
if len(line) >= 56 {
d1 = append(d1, line[:53])
d2 = append(d2, line[56:])
} else {
d1 = append(d1, line)
d2 = append(d2, line)

fmt.Println("56 ?/", len(line))
}
}

diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
A:        difflib.SplitLines(strings.Join(d1, "\n")),
B:        difflib.SplitLines(strings.Join(d2, "\n")),
Context:  3,
})

f, _ := os.Create("f0.png")
defer f.Close()
f1, _ := os.Create("f1.png")
defer f1.Close()
f2, _ := os.Create("f2.png")
defer f2.Close()

//fmt.Println(diff)

for i, line := range strings.Split(diff, "\n") {
if len(line) == 0 {
continue
}

var bs []byte
if !strings.HasPrefix(line, "@") {//&& (strings.HasPrefix(line, "+") || strings.HasPrefix(line, "-")) {
hexData := line[1:]
fmt.Println(i, line)
bs, _ = hexStringToBytes(hexData)
//hexData := strings.TrimSpace(line[1:])
//bs, _ = hexStringToBytes(hexData)
}
//if len(bs) < 1 { continue } 
//switch {
if strings.HasPrefix(line, "+") {
f1.Write(bs)
//fmt.Println("+/", line)
} else if strings.HasPrefix(line, "-") {
f2.Write(bs)
//fmt.Println("-/", strconv.Quote(line))
} else if strings.HasPrefix(line, " "){
f.Write(bs)
fmt.Println(i, "prefix!=@/", len(bs), line)
} 
/*else {
fmt.Println(i, "default/", line)
}*/
//default:
//fmt.Println(i, "default/", line)
//}
}
}

func hexStringToBytes(hexStr string) ([]byte, error) {
hexStr = strings.ReplaceAll(hexStr, " ", "")
if len(hexStr)%2 != 0 {
hexStr = "0" + hexStr
}
return hex.DecodeString(hexStr)
}

