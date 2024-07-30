package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

var PAGE string

func main(){
    fmt.Println(PAGE, "\nbody ends/")
}


func init(){

    URL := "http://www.pythonchallenge.com/pc/return/italy.html"
    u, p := "huge","file"

    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(u,p)
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    PAGE = string(body)
}

