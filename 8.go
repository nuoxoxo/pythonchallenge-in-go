package main

import (
    "fmt"
    "bytes"
    "compress/bzip2"
    "io/ioutil"
    "net/http"
    "regexp"
    "strconv"
)

var un, pw string

func main () {
    User := []byte(un)
    readerUser := bzip2.NewReader(bytes.NewReader( User ))
    bytesUser, _ := ioutil.ReadAll(readerUser)
    stringUser := string( bytesUser )
    fmt.Println("usr/", stringUser)

    Pass := []byte(pw)
    readerPass := bzip2.NewReader(bytes.NewReader( Pass ))
    bytesPass, _ := ioutil.ReadAll(readerPass)
    stringPass := string( bytesPass )
    fmt.Println("pwd/", stringPass)
}

func init(){

    URL := "http://www.pythonchallenge.com/pc/def/integrity.html"
    resp, _ := http.Get(URL)
    body, _ := ioutil.ReadAll( resp.Body )
    defer resp.Body.Close()

    re := regexp.MustCompile(`(?s)'(.*?)'`)
    matches := re.FindAllStringSubmatch(string(body), -1)
    for i, m := range matches {
        for j, m2 := range m {
            fmt.Println(i, j, m2)
        }
    }

    // Trick/
    //  quotes/ adde double quotes to enable Go-style string literal
    //  unquote/ 'decode' all escape sequences like this
    un, _ = strconv.Unquote("\"" + matches[0][1] + "\"")
    pw, _ = strconv.Unquote("\"" + matches[1][1] + "\"")
    fmt.Println("u/", un)
    fmt.Println("p/", pw)
}
