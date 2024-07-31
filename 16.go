package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "image"
    "image/gif"
    "image/color"
    "os"
    "reflect"
)

func main(){
    FILE, _ := os.Open("files/mozart.gif")
    defer FILE.Close()
    mozart, _ := gif.Decode(FILE)
    bounds := mozart.Bounds()
    var R, C, r, c int
    R, C = bounds.Dy(), bounds.Dx()
    fmt.Println("size/", bounds, reflect.TypeOf(bounds))
    fmt.Println("rows/", R, "cols/", C)
    res := image.NewPaletted( bounds, nil )
    fmt.Println("init/", res, reflect.TypeOf(res))
    r = 0
    for r < R {
        // row reprented as a rgb(a) slice
        row := make([]color.Color, C)
        c = 0
        for c < C {
            cl := mozart.At(c, r)
            fmt.Println("color/", cl, reflect.TypeOf(cl))
            row[c] = cl
        }
        // here we do longest uni-char substring 
        // todo...
        r++
    }

}

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

}

