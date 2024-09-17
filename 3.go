package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

func main(){
    URL := "http://www.pythonchallenge.com/pc/def/equality.html"
    response, _ := http.Get(URL)
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    re := regexp.MustCompile("[^A-Z]+[A-Z]{3}([a-z])[A-Z]{3}[^A-Z]+")
    matches := re.FindAllStringSubmatch( string(body), -1 )
    res := ""
    for i, pair := range matches {
        res += pair[1]
        fmt.Println(string(i + '0') + "/", pair[1])
    }
    fmt.Println("res/", res)
}

