package main

import (
    "fmt"
    "regexp"
    "archive/zip"
    "io/ioutil"
    "os"
    "net/http"
    "io"
)

func main(){

    // zip reader
    openreader, err := zip.OpenReader("./channel.zip")
    if err != nil {
        fmt.Println("open reader/", err)
    }
    defer openreader.Close()

    // zip files - single file reader
    // begin w/ 90052
    re := regexp.MustCompile(`nothing is (\d+)`)
    comments := []string{}
    s := "90052"
    e := ".txt"
    for {
        // get access to a file in zip
        fname := s + e
        f, err := read_inside_zip(openreader, fname)
        if err != nil {
            fmt.Println("read inside zip/", fname, err)
        }
        // open the target file
        fopen, err := f.Open()
        if err != nil {
            fmt.Println("f.open()/", fname, err)
        }
        // get its content & comment
        cont, err := ioutil.ReadAll(fopen)
        cmt := f.Comment
        char := cmt
        if char == "\n" { char = "(newline)" }
        if char == " " { char = "(space)" }
        fmt.Println(s, "\t", string(cont), "\t", char)
        comments = append(comments, cmt)
        // get the next filename
        match := re.FindSubmatch(cont)
        if match == nil {
            break
        }
        //fmt.Println("match[0]/", string(match[0]))
        //fmt.Println("match[1]/", string(match[1]))
        s = string(match[1])
    }
    res := ""
    res2 := ""
    seen := make(map[string]bool)
    for _, cmt := range comments {
        res += cmt
        if cmt != "\n" && !seen[cmt] {
            seen[cmt] = true
            res2 += cmt
        }
    }
    fmt.Println(res)
    fmt.Println(res2)
}

// unzip

func read_inside_zip (openreader * zip.ReadCloser, target string) (*zip.File, error) {

    for _, f := range openreader.File {
        if f.Name != target {
            continue
        }
        return f, nil
    }
    return nil, fmt.Errorf("err/not found: %s", target)
}

// download channel.zip
func init(){

    URL := "http://www.pythonchallenge.com/pc/def/channel.zip"
    f, err := os.Create("channel.zip")
    fmt.Println("err/download", err)
    defer f.Close()

    resp, err := http.Get( URL )
    fmt.Println("err/GET", err, resp.StatusCode)
    defer resp.Body.Close()

    _, _ = io.Copy(f, resp.Body)


}


