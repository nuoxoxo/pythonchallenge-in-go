package main

import (
    "fmt"
    "reflect"
    "os"
    "io/ioutil"
    "strconv"
    "net/http"
)

func main(){
    content, _ := ioutil.ReadFile("files/evil2.gfx")
    fmt.Println(content[:42], reflect.TypeOf(content) ) // uint8
    N := len(content)
    i := 0
    for i < 5 {
        name := "files/0" + strconv.Itoa(i) + ".jpg"
        outfile, _ := os.Create( name )
        j := i
        for j < N {
            outfile.Write( content[j : j + 1] )
            j += 5
        }
        outfile.Close()
        i++
    }

    // extra : get a msg for next level
    URL := "http://www.pythonchallenge.com/pc/return/evil4.jpg"
    user := "huge"
    pass := "file"
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}
    req.SetBasicAuth( user, pass )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("msg/", string(body))
}



