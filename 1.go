package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

func maketrans(s string) string {
    res := []rune{}
    for _, char := range s {
        nxt := int(char)
        if 'a' <= char && char <= 'z' {
            nxt += 2
            if nxt >= 'z' {
                nxt = nxt - 'z' + 'a' - 1
            }
        }
        res = append(res, rune(nxt))
    }
    return string(res)
}

func main() {
    resp, _ := http.Get("http://www.pythonchallenge.com/pc/def/map.html")
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    re := regexp.MustCompile(`(?s)<font color="#f000f0">(.*?)</tr></td>`)
    sub := re.FindAllStringSubmatch(string(body), -1)[0][1]
    fmt.Println("sub/", sub)

    fmt.Println("str/", maketrans(sub))
    fmt.Println("map/", maketrans("map"))
}
