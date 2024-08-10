package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    _"reflect"
    "regexp"
)


func main(){

    // lv.27
    data, _ := getbody("speedboat.html")
    fmt.Println(string(data))
    yell("body ends/\n")
    re := regexp.MustCompile(`(?s)<img src="(.*?)"`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    // zigzag
    data, _ = getbody(sub)
    fmt.Println("data/snippet", data[:42], "len/", len(data))
    fmt.Println("div3/", len(data) / 3, "mod3/", len(data) % 3)

}

func getbody(sub string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/hex/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) { fmt.Println(Yell + s + Rest) }
func cy(s string)   { fmt.Println(Cyan + s + Rest) }




