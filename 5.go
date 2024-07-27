package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/gob"
    "bytes"
)

type CharCount struct {
    Char    string
    Count   int
}

func main(){

    counters := [][]CharCount{}

    URL := "http://www.pythonchallenge.com/pc/def/banner.p"
    resp, _ := http.Get(URL)
    body, _ := ioutil.ReadAll(resp.Body)
    decoder := gob.NewDecoder( bytes.NewReader(body) )
    err := decoder.Decode( &counters )
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, counter := range counters {
        line := ""
        for _, pc := range counter {
            i := 0
            for i < pc.Count {
                line += pc.Char
                i++
            }
        }
        fmt.Println( line )
    }
    fmt.Println("line/")
}


