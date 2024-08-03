package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    _"strings"
    "regexp"
    "strings"
    "encoding/base64"
)

var URL, BODY, filename, b64 string

func main(){

    // parsing done in init
    b64data, _ := base64.StdEncoding.DecodeString( b64 )
    ioutil.WriteFile(filename, b64data, 0644)
}

func init(){

    URL = "http://www.pythonchallenge.com/pc/hex/bin.html"
    ups, sep := "butterfly", 6
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(ups[: sep], ups[sep :])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body), "\nbody ends/\n\n")
    BODY = string(body)

    // filename
    re := regexp.MustCompile(`(?s)name="(.*?)"`)
    matches := re.FindAllStringSubmatch(BODY, -1)
    filename = matches[0][1]

    // get the bound
    re = regexp.MustCompile(`(?s)boundary="(.*?)"`)
    matches = re.FindAllStringSubmatch(BODY, -1)
    bound := "--" + matches[0][1]
    N := len(matches[0])
    fmt.Println("len/", len(matches[0]), "matches/", matches)
    fmt.Println("match/", strconv.Quote( matches[0][1] ))
    fmt.Println("modf./", strconv.Quote( bound ))
    fmt.Println("check/", "--===============1295515792==")

    // get the bounded trunk which should look base64 encoded
    re = regexp.MustCompile(fmt.Sprintf(`(?s)%s(.*?)%s`, bound, bound))
    matches = re.FindAllStringSubmatch(BODY, -1)
    offset := 42
    N = len(matches[0][1])
    end := N - offset
    fmt.Println("\nlen/", len(matches[0]))
    fmt.Println("aft/", matches[0][1][: offset], "bef/", matches[0][1][end :])

    trunk := strings.Split(matches[0][1], "\n\n")
    b64 = trunk[1]
    N = len( b64 )
    end = N - offset
    fmt.Println("len/", len(trunk))
    fmt.Println("aft/", trunk[1][: offset], "bef/", trunk[1][end :])
}

