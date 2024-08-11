package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "bytes"
    "compress/bzip2"
)

func main(){
    page, _ := getbody("guido.html")
    lines := strings.Split(strings.ReplaceAll(string(page), "\r\n", "\n"), "\n")
    lens := []uint8{}
    for i, line := range lines {
        // parsing empty lines
        if strings.TrimSpace(line) == "" {
            if len(line) > 255 { panic("tofu/") } 
            fmt.Println(i, len(line))
            lens = append(lens, uint8(len(line)))
        }
    }
    buff := bytes.NewBuffer(lens)
    reader := bzip2.NewReader(buff)
    res, err := ioutil.ReadAll(reader)
    fmt.Println(yell("err/readall"), err)
    fmt.Println("res/", string(res))
}

func getbody(sub string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/ring/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "repeat", "switch" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(data), yell("\bbody ends/"))
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) string { return Yell + s + Rest }
func cyan(s string) string { return Cyan + s + Rest }


