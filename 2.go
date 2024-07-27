package main

import (
    "fmt"
    "io/ioutil"
    _"log"
    "net/http"
    "regexp"
)

// Compile(expr string)
// MustCompile(expr string) - is like Compile but panics if an expr cannot be parsed

func main(){

    URL := "http://www.pythonchallenge.com/pc/def/ocr.html"
    response, _ := http.Get( URL )
    defer response.Body.Close()
    body, _ := ioutil.ReadAll( response.Body )
    //fmt.Println(string(body), "\nend/")
    //re, _ := regexp.Compile(`<!--(.*?)-->`) 

    // (?s) is a flag for `dotall` mode ie. single line incl. newlines
    re := regexp.MustCompile(`(?s)<!--(.*?)-->`)
    
    //matches := re.FindAllString( string(body) , -1)
    matches := re.FindAllStringSubmatch(string(body), -1)
    fmt.Println("len/", len(matches))
    /*
    for i, match := range matches {
        for j, m := range match {
            fmt.Println(i, j, m[:10])
        }
    }
    */
    // it's the 2nd match
    instr, cmt := matches[0][1], matches[1][1]
    fmt.Println(cmt, instr)
    dict := make(map[rune]int)
    res := []rune{}
    for _, c := range (cmt) {
        if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
            dict[c]++
            res = append(res, c)
        }
    }
    for k, v := range dict {
        fmt.Printf("%c/ %d \n", rune(k), v)
    }
    fmt.Println("res/", string(res))
}
