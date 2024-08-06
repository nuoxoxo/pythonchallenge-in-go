package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    _"reflect"
    _"strconv"
)

const URL string = "http://www.pythonchallenge.com/pc/hex/"
const UPS string = "butterfly"

func main(){

    // todo
}



func init(){

    // Step - go to the url
    sub1 := URL + "idiot2.html"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", sub1, nil)

    req.SetBasicAuth(UPS[:6], UPS[6:])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    BODY := string(body)
    // fmt.Println(BODY, "body/ends")

    // Step - go get `unreal`
    re := regexp.MustCompile(`(?s)src="(.*?)"`)
    matches := re.FindAllStringSubmatch(BODY, -1)
    path := matches[0][1]
    sub2 := URL + path // ie. ___.jpg

    conn = & http.Client{}
    req, _ = http.NewRequest("GET", sub2, nil)

    req.SetBasicAuth(UPS[:6], UPS[6:])
    resp, _ = conn.Do(req)
    defer resp.Body.Close()

    body, _ = ioutil.ReadAll(resp.Body)
    data := []byte(string(body))

    // fmt.Println("header/gross", resp.Header, reflect.TypeOf(resp.Header))
    for k, v := range resp.Header { fmt.Println("head/", k, v) }
    fmt.Println("sizeof/", len(data))

    // observation/
    //  Content-Range [bytes 0-30202/2123456789]
    //  Content-Length [30203]
    // sizeof/ 30203 - should be == content-length
    // conclusion/ read again w/ range set

    // Step - set range in header and go again
    conn = & http.Client{}
    req, _ = http.NewRequest("GET", sub2, nil)

    req.Header.Set("Range", "bytes=" + resp.Header["Content-Length"][0] + "-")

    req.SetBasicAuth(UPS[:6], UPS[6:])
    resp, _ = conn.Do(req)
    defer resp.Body.Close()

    body, _ = ioutil.ReadAll(resp.Body)
    fmt.Println("\nbody/", string(body))
    for k, v := range resp.Header { fmt.Println("head/", k, v) }
}


