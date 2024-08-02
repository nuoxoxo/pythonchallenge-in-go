package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    _"strings"
    "regexp"
    "os"
    "io"
)

//var URL, BODY string


func main(){





}

func init(){

    URL := "http://www.pythonchallenge.com/pc/hex/bin.html"
    ups, sep := "butterfly", 6
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(ups[: sep], ups[sep :])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body), "\nbody ends/\n\n")
    BODY := string(body)

    re := regexp.MustCompile(`(?s)name="(.*?)"`)
    matches := re.FindAllStringSubmatch(BODY, -1)

    fmt.Println("len/", len(matches), "matches/", matches)

    filename := matches[0][1]
    fmt.Println("match/", strconv.Quote( filename ))

    conn = & http.Client{}
    req, err = http.NewRequest("GET", URL + "/" + filename, nil)
    
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(ups[: sep], ups[sep :])
    resp, _ = conn.Do(req)
    defer resp.Body.Close()
    fmt.Println("status/", resp.StatusCode, resp.Status)

    f, _ := os.Create(filename)
    defer f.Close()
    _, _ = io.Copy(f, resp.Body)
}

