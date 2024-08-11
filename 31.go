package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
)

func main(){}

func init(){
    // main page
    data, _ := getbody("ring/grandpa.html", "repeat", "switch")
    fmt.Println(string(data))
    yell("body ends/\n")
    // href
    re := regexp.MustCompile(`(?s)<a href="../(.*?)"`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(cyan("sub/"), sub)
    data, _ = getbody(sub, "kohsamui", "thailand")
    fmt.Println(string(data))
    yell("body ends/\n")
    // mandelbrot
    prev := sub[:5]
    re = regexp.MustCompile(`(?s)img src="(.*?)"`)
    sub = re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(cyan("prev/sub"), prev, sub)
    data, _ = getbody(prev + sub, "kohsamui", "thailand")
    fmt.Println(string(data)[:123])
    yell("body ends/\n")
}

func getbody(sub, u, p string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth(u, p)
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) string { return Yell + s + Rest }
func cyan(s string) string { return Cyan + s + Rest }

