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
)

var PAGE string

func main(){
    fmt.Println(PAGE, "\nbody ends/\n\n")

    raw, _ := os.Open("files/wire.png")
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

    f, _ := os.Create("files/res.jpg")
    defer f.Close()
    jpeg.Encode(f, res, nil)
    fmt.Println("res/", reflect.TypeOf(res))
    fmt.Println("out/", f, reflect.TypeOf(f))
}

func init(){
    URL := "http://www.pythonchallenge.com/pc/return/italy.html"
    u, p := "huge","file"

    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(u,p)
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    PAGE = string(body)
}


