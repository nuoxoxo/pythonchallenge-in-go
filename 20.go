package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "strings"
    "strconv"
    _"reflect"
)

const URL string = "http://www.pythonchallenge.com/pc/hex/"
const UPS string = "butterfly"

func main(){

    // todo
}

func iter2(url, user, pass string, start int){

    i := 0
    for {
        if i==200 {break}
        fmt.Println("\ninside iter1/", i)
        conn := & http.Client{}
        req, err := http.NewRequest("GET", url, nil)
        fmt.Println("req/err", err)
        req.Header.Set("Range", "bytes=" + strconv.Itoa(start) + "-")
        req.SetBasicAuth(user, pass)
        resp, err := conn.Do(req)
        fmt.Println("resp/err", err)
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Println("body/err", err)
        fmt.Println("\nbody/", string(body))
        for k, v := range resp.Header { fmt.Println("head/", k, v) }

        // break
        if resp.StatusCode != 200 && resp.StatusCode != 206 || 
            len(resp.Header["Content-Range"][0]) == 0 {
            fmt.Println("break/status code", resp.StatusCode)
            fmt.Println("break/status text", http.StatusText(resp.StatusCode))
        }
        start--
        i++
    }
    /*
    var rev string
    b := string(body)
    i = len(b) - 1
    for i > -1 {
        rev += string(b[i])
        i--
    }
    fmt.Println("res/", rev)
    */
}

func iter1(url, user, pass, start string){

    i := 0
    for {
        fmt.Println("\ninside iter1/", i)
        conn := & http.Client{}
        req, err := http.NewRequest("GET", url, nil)
        fmt.Println("req/err", err)
        req.Header.Set("Range", "bytes=" + start + "-")// + end)
        req.SetBasicAuth(user, pass)
        resp, err := conn.Do(req)
        fmt.Println("resp/err", err)
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Println("body/err", err)
        fmt.Println("\nbody/", string(body))
        for k, v := range resp.Header { fmt.Println("head/", k, v) }

        // break
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

func init(){

    // Step - go to the url
    sub1 := "idiot2.html"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub1, nil)

    req.SetBasicAuth(UPS[:6], UPS[6:])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    BODY := string(body)
    // fmt.Println(BODY, "body/ends")

    // Step - go get `unreal.jpg`
    re := regexp.MustCompile(`(?s)src="(.*?)"`)
    matches := re.FindAllStringSubmatch(BODY, -1)
    sub2 := matches[0][1]

    conn = & http.Client{}
    req, _ = http.NewRequest("GET", URL + sub2, nil)

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

    // Step - get len and range for iter fns
    ContentLength := resp.Header["Content-Length"][0]
    cr := resp.Header["Content-Range"][0]
    idx := strings.Index(cr, "/")
    ContentEnd, _ := strconv.Atoi(cr[idx + 1:])

    // Step - set range in header and go again
    conn = & http.Client{}
    req, _ = http.NewRequest("GET", URL + sub2, nil)

    req.Header.Set("Range", "bytes=" + ContentLength + "-")

    req.SetBasicAuth(UPS[:6], UPS[6:])
    resp, _ = conn.Do(req)
    defer resp.Body.Close()

    body, _ = ioutil.ReadAll(resp.Body)
    fmt.Println("\nbody/", string(body))
    for k, v := range resp.Header { fmt.Println("head/", k, v) }

    // Step - iter
    //  iter1 ends at 5th iteration
    iter1( URL + sub2, UPS[:6], UPS[6:], ContentLength )
    iter2( URL + sub2, UPS[:6], UPS[6:], ContentEnd )

}


