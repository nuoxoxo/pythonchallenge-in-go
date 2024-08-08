package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

const URL string = "http://www.pythonchallenge.com/pc/hex/"
//const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"

func main(){}

func init(){

    // GET
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + "ambiguity.html", nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    temp, _ := ioutil.ReadAll(resp.Body)
    body := string(temp)

    fmt.Println(body, "\nbody ends/")
}


