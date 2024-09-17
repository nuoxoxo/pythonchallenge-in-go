package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

// re
//  Compile(expr string)
//  MustCompile(expr string) - like Compile but panic if expr cannot be parsed

func main(){

    URL := "http://www.pythonchallenge.com/pc/def/ocr.html"
    response, _ := http.Get( URL )
    defer response.Body.Close()
    body, _ := ioutil.ReadAll( response.Body )
    re := regexp.MustCompile(`(?s)<!--(.*?)-->`)
    matches := re.FindAllStringSubmatch(string(body), -1)

    fmt.Println("len/", len(matches))

    instr, cipher := matches[0][1], matches[1][1]

    fmt.Println("\ncipher/", cipher[:42 * 10])
    fmt.Println("\nmanual/", instr)

    dict := make(map[rune]int)
    res := []rune{}
    for _, c := range (cipher) {
        if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
            res = append(res, c)
            dict[c]++
        }
    }
    for k, v := range dict {
        fmt.Printf("%c/ %d \n", rune(k), v)
    }
    fmt.Println("res/", string(res))
}
