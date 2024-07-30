package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main(){
    
}

func init(){
    URL := "http://www.pythonchallenge.com/pc/return/mozart.html"
    ups, mid := "hugefile", 4
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(ups[:mid], ups[mid:])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body), "\nbody ends/\n\n")

}

