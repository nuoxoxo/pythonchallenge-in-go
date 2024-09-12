package main

import (
    "fmt"
    "net/http"
    "github.com/hydrogen18/stalecucumber"
)

func main() {
    resp, err := http.Get("http://www.pythonchallenge.com/pc/def/banner.p")
    if err != nil {
        fmt.Println("resp/err", err)
    }
    defer resp.Body.Close()

    // some stalecucumber magic
    unpickled, err := stalecucumber.Unpickle(resp.Body)
    if err != nil {
        fmt.Println("stale/err", err)
    }
    data, ok := unpickled.([]interface{})
    if !ok {
        fmt.Println("assert/data", "not []interface{}")
    }

    // loop through line by line
    for _, d := range data {
        sub, ok := d.([]interface{})
        if !ok {
            fmt.Println("assert/sub", "not []interface{}")
        }

        line := ""
        for _, pair := range sub {
            kv, ok := pair.([]interface{})
            if !ok {
                fmt.Println("assert/pair", "not []interface{}")
            }
            if len(kv) != 2 {
                fmt.Println("assert/err", "not key-value pair")
            }
            k, ok := kv[0].(string)
            if !ok {
                fmt.Println("assert/not string")
            }
            v, ok := kv[1].(int64)
            if !ok {
                fmt.Println("assert/not int64 - has to be int64")
            }
            i := 0
            for i < int(v) {
                line += k
                i++
            }
        }
        fmt.Println("line/ ", line)
    }
}

