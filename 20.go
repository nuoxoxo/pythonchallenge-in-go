package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "strings"
    "strconv"
    "os"
    _"reflect"
)

const URL string = "http://www.pythonchallenge.com/pc/hex/"
const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"

// observation/
//  Content-Range [bytes 0-30202/2123456789]
//  Content-Length [30203]
// sizeof/ 30203 - should be == content-length
// conclusion/ read again w/ range set

func main(){

    // go to the URL
    resp, _ := getresponse("idiot2.html", "")
    defer resp.Body.Close()
    bodybytes, _ := ioutil.ReadAll(resp.Body)
    body := string(bodybytes)
    fmt.Println(body, "body/ends")

    // get path - `unreal.jpg`
    re := regexp.MustCompile(`(?s)src="(.*?)"`)
    matches := re.FindAllStringSubmatch(body, -1)
    sub := matches[0][1]
    resp, _ = getresponse(sub, "")
    defer resp.Body.Close()

    // get len and range for iter
    ContentLength := resp.Header["Content-Length"][0]
    cr := resp.Header["Content-Range"][0]
    idx := strings.Index(cr, "/")
    ContentEnd, _ := strconv.Atoi(cr[idx + 1:])
    resp, _ = getresponse(sub, ContentLength)
    defer resp.Body.Close()
    bodybytes, _ = ioutil.ReadAll(resp.Body)
    fmt.Println("\nbody/", string(bodybytes))
    for k, v := range resp.Header { fmt.Println("head/", k, v) }

    iter1( sub, ContentLength )
    iter2( sub, ContentEnd )
}

func iter2(sub string, start int){

    i := 0
    found := false
    for {
        fmt.Println("\ninside iter2/", i)
        resp, _ := getresponse(sub, strconv.Itoa(start))
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        if resp.StatusCode != 200 && resp.StatusCode != 206 || 
            len(resp.Header["Content-Range"][0]) == 0 {
            fmt.Println("break/status code", resp.StatusCode)
            fmt.Println("break/status text", http.StatusText(resp.StatusCode))
            return
        }

        if found {
            for k, v := range resp.Header { fmt.Println("head/", k, v) }
            _ = os.WriteFile( "readme.txt" , body, 0644)
            return
            // this part to be reproduced in puzzle 21
        }
        s := strings.TrimSpace(string(body))
        if ! strings.Contains(s, "hiding") {
            fmt.Println(i, "original/", s)
            fmt.Println(i, Cyan + "reversed/", strrev(s), Rest)
            start--
        } else {
            at := strings.Index(s, "at")
            pos, _ := strconv.Atoi(s[at + 3 : len(s) - 1])
            start = pos
            found = true
        }
        i++
    }
}

func strrev (s string) string {
    var res string
    for i:= len(s) - 1; i > -1; i-- {res += string(s[i])}
    return res
}

func iter1(sub, start string){

    i := 0
    for {
        fmt.Println("\ninside iter1/", i)
        resp, _ := getresponse(sub, start)
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Println("body/err", err)
        if i == 4 { fmt.Print(Yell) }
        fmt.Println("\nbody/", string(body))
        if i == 4 { fmt.Print(Rest) }
        for k, v := range resp.Header { fmt.Println("head/", k, v) }

        if resp.StatusCode != 200 && resp.StatusCode != 206 || 
            len(resp.Header["Content-Range"][0]) == 0 {
            fmt.Println("break/status code", resp.StatusCode)
            fmt.Println("break/status text", http.StatusText(resp.StatusCode))
            break
        }
        // parse header
        re := regexp.MustCompile(`(?s)bytes (.*?)/`)
        matches := re.FindAllStringSubmatch(resp.Header["Content-Range"][0], -1)
        match := strings.Split(matches[0][1], "-")
        fmt.Println("split/", match)
        temp, _ := strconv.Atoi(match[1])
        start = strconv.Itoa(temp + 1)
        fmt.Println("start/", start)
        i++
    }
}

func getresponse(sub, rangestart string) (*http.Response, error) {
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "butter", "fly" )
    if rangestart != "" {
        req.Header.Set("Range", "bytes=" + rangestart + "-")// + end)
    }
    response, _ := conn.Do(req)
    return response, nil
    // todo/ defer response.Body.Close() after called
}


