package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    _"strings"
    _"strconv"
    _"os"
    _"reflect"
)

const URL string = "http://www.pythonchallenge.com/pc/hex/"
const Yell string = "\033[33m" 
const Cyan string = "\033[36m" 
const Rest string = "\033[0m"

func main(){}

func init(){

    // GET
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + "copper.html", nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    // to be told how url should be modified
    temp, _ := ioutil.ReadAll(resp.Body)
    body := string(temp)
    re := regexp.MustCompile(`(?s)maybe (.*?) would`)
    matches := re.FindAllStringSubmatch(body, -1)
    sub2 := matches[0][1] // white.gif

    conn = & http.Client{}
    req, _ = http.NewRequest("GET", URL + sub2, nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ = conn.Do(req)
    defer resp.Body.Close()
    temp, _ = ioutil.ReadAll(resp.Body)
    body = string(temp)
    fmt.Println(body[:42 * 42], "\nbody ends/")
}


