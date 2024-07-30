package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "reflect"
    "regexp"
    "strings"
    _"strconv"
    "net/url"
    "compress/bzip2"
    "io"
    "github.com/kolo/xmlrpc"
)

const endl string = "\n"
const linkedlist_passed bool = true//false
var linkedlist_res string
//var phonebook_res string

func main(){

    // back to linkedlist - too long, use a helper func to bypass the loop
    if !linkedlist_passed {
        linkedlist_res = linkedlist_revisited()
    }

    fmt.Println("stuff/", linkedlist_res)
    decoded, _ := url.QueryUnescape( linkedlist_res )
    fmt.Println("dcode/", decoded[:42], "len/", len(decoded))
    reader := strings.NewReader(decoded)
    bzipreader := bzip2.NewReader(reader)
    decompressed := new(strings.Builder)
    io.Copy(decompressed, bzipreader)
    fmt.Println("dcomp/", decompressed, endl)

    // back to phonebook
    URL := "http://www.pythonchallenge.com/pc/phonebook.php"
    conn, err := xmlrpc.NewClient(URL, nil)
    if err != nil {
        fmt.Println("xmlrpc.NewClient/", err)
    }
    var res interface{}
    err = conn.Call("phone", "Leopold", &res)
    fmt.Println("Leo/", res, "type/", reflect.TypeOf(res), endl)

    // bring cookie as msg
    dcomp := decompressed.String()//string(decompressed)
    l := strings.Index(dcomp, "\"") + 1
    fmt.Println("len/", len(dcomp), "l/", l)
    r := strings.Index(dcomp[l:], "\"") + l
    fmt.Println("len/", len(dcomp), "r/", r)
    MSG := dcomp[l : r]
    fmt.Println("msg/", MSG, endl)

    URL = "http://www.pythonchallenge.com/pc/stuff/violin.php"
    req, _ := http.NewRequest("GET", URL, nil)
    req.Header.Set("Cookie", "info=" + MSG)
    clt := & http.Client{}
    response, _ := clt.Do( req )
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body), "\nbody ends/\n\n")
}

func linkedlist_revisited() string {

    // back to linkedlist
    LL := "http://www.pythonchallenge.com/pc/def/linkedlist.php"
    ll, _ := http.Get(LL)
    defer ll.Body.Close()
    for _, ckie := range ll.Cookies() {
        fmt.Println("cookie/", ckie, "\ntype/", reflect.TypeOf(ckie))
    }

    // follow the hint
    TAIL := "12345"//"96070"//"12345"
    busyURL := "http://www.pythonchallenge.com/pc/def/linkedlist.php?busynothing="
    re := regexp.MustCompile(`and the next busynothing is (.*)`)
    res := ""
    count := 0
    for {
        response, _ := http.Get(busyURL + TAIL)
        body, _ := ioutil.ReadAll(response.Body)
        fmt.Println(count, "\b/", string(body))
        count++
        matches := re.FindAllStringSubmatch(string(body), -1)
        if matches == nil {
            for _, ckie := range response.Cookies() {
                fmt.Println("busy cookie/break", ckie, "\ntype/", reflect.TypeOf(ckie))
            }
            break
        } else {
            TAIL = matches[0][1]
            for _, ckie := range response.Cookies() {
                fmt.Println("busy cookie/", ckie, "\ntype/", reflect.TypeOf(ckie))
                fmt.Println("cookie value/", string(ckie.Value), reflect.TypeOf(ckie))
                //fmt.Println("re2/", matches2, ckie.Value)
                res += string(ckie.Value)
            }
        }
    }
    return res
}

//

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

    // skip what's done
    linkedlist_res = "BZh91AY%26SY%94%3A%E2I%00%00%21%19%80P%81%11%00%AFg%9E%A0%20%00hE%3DM%B5%23%D0%D4%D1%E2%8D%06%A9%FA%26S%D4%D3%21%A1%EAi7h%9B%9A%2B%BF%60%22%C5WX%E1%ADL%80%E8V%3C%C6%A8%DBH%2632%18%A8x%01%08%21%8DS%0B%C8%AF%96KO%CA2%B0%F1%BD%1Du%A0%86%05%92s%B0%92%C4Bc%F1w%24S%85%09%09C%AE%24"
}

