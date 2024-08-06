package main

import (
    "fmt"
    "reflect"
    "os"
    "io"
    "io/ioutil"
    "strconv"
    "net/http"
)

func main(){

    Filename := "evil2.gfx"

    f, err := os.Create( Filename )
    defer f.Close()

    URL := "http://www.pythonchallenge.com/pc/return/" + Filename
    ups, mid := "hugefile", 4
    user, pass := ups[:mid], ups[mid:]
    conn := &http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    fmt.Println("err/GET", err)

    req.SetBasicAuth( user, pass )
    resp, err := conn.Do( req )
    fmt.Println("err/conn.DO", err, resp.StatusCode)
    defer resp.Body.Close()

    _, _ = io.Copy(f, resp.Body)

    content, _ := ioutil.ReadFile( Filename )
    fmt.Println(content[:42], reflect.TypeOf(content) ) // uint8
    N := len(content)
    i := 0
    for i < 5 {
        name := "12_" + strconv.Itoa(i) + ".jpg"
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
    URL = "http://www.pythonchallenge.com/pc/return/evil4.jpg"
    conn = & http.Client{}
    req, err = http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}
    req.SetBasicAuth( user, pass )
    resp, _ = conn.Do( req )
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("msg/", string(body))
}



