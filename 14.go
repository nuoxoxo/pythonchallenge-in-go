package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "reflect"
    "io"
)

var PAGE string
var Filename string = "wire.png"

func main(){
    fmt.Println(PAGE, "\nbody ends/\n\n")

    raw, err := os.Open( Filename )
    fmt.Println("err/open", err)
    defer raw.Close()
    var wire image.Image 
    wire, _ = png.Decode(raw)

    // let's draw
    res := image.NewRGBA(image.Rect(0,0,100,100))
    sx, ex := 0, 100
    sy, ey := 0, 100
    var x, y int
    N := 10000
    w := 0

    var valid func(int, int, int) bool
    valid = func(x,y,n int) bool { return -1<x && x<100 && -1<y && y<100 && n<10000 }

    for w < N && sx < ex && sy < ey {
        // fmt.Println("\ndbg/", sx, ex, sy, ey, "w/", w)
        // >
        x, y = sx, sy
        for x < ex && valid(x,y,w) {
            res.Set(x, y, wire.At(w, 0))
            x++
            w++
        }
        sy++

        // v
        x, y = ex - 1, sy
        for y < ey && valid(x,y,w) {
            res.Set(x, y, wire.At(w, 0))
            y++
            w++
        }
        ex--

        // <
        x, y = ex - 1, ey - 1
        for x > sx - 1 && valid(x,y,w) {
            res.Set(x, y, wire.At(w, 0))
            x--
            w++
        }
        ey--

        // ^
        x, y = sx, ey - 1
        for y > sy - 1 && valid(x,y,w) {
            res.Set(x, y, wire.At(w, 0))
            y--
            w++
        }
        sx++
        // fmt.Println("dbg2/", sx, ex, sy, ey, "w/", w)

    }

    f, _ := os.Create("res14.jpg")
    defer f.Close()
    jpeg.Encode(f, res, nil)
    fmt.Println("res/", reflect.TypeOf(res))
    fmt.Println("out/", f, reflect.TypeOf(f))
}

func init(){

    // show bert
    URL := "http://www.pythonchallenge.com/pc/return/italy.html"
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    fmt.Println("err GET/url", err)

    req.SetBasicAuth("huge", "file")
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    PAGE = string(body)

    // todo/ download wire
    URL = "http://www.pythonchallenge.com/pc/return/" + Filename
    conn = & http.Client{}
    req, err = http.NewRequest("GET", URL, nil)
    fmt.Println("err GET/jpg", err)
    req.SetBasicAuth("huge", "file")
    resp, _ = conn.Do(req)
    defer resp.Body.Close()

    f, err := os.Create( Filename )
    defer f.Close()
    _, _ = io.Copy(f, resp.Body)
}


