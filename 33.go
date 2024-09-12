package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "reflect"
)

func main() {
    data, _ := getbody("rock/beer.html", "kohsamui", "thailand")
    fmt.Println(string(data))

    re := regexp.MustCompile(`(?s)<img src="(.*?)"`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    data, _ = getbody("rock/" + sub, "kohsamui", "thailand")
    fmt.Println(Cyan + "sub/data on newline" + Rest)
    fmt.Println(string(data[:512]))
    fmt.Println(reflect.TypeOf(data))

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
func yell(s string) { fmt.Println( Yell + s + Rest )}
func cyan(s string) { fmt.Println( Cyan + s + Rest )}



