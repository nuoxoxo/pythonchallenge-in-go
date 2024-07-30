package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "image"
    "image/jpeg"
    "image/png"
    _"image/color"
    "os"
)

var PAGE string

func main(){
    fmt.Println(PAGE, "\nbody ends/\n\n")

    // let's draw
    raw, _ := os.Open("files/wire.png")
    defer raw.Close()
    var wire image.Image 
    wire, _ = png.Decode(raw)
    res := image.NewRGBA(image.Rect(0,0,101,101))
    D := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    i := 0
    offset := 100
    c, r := 0, 0
    for i < 100 {
        color := wire.At(i, 0)//.RGBA()
         
        for _, d := range D {
            r, c = r + d[0] * offset, c + d[1] * offset
            res.Set(c, r, color)
            fmt.Println(i, c, r)
            i += 1
            offset -= 1
        }
    }
    f, _ := os.Create("files/res.jpg")
    defer f.Close()
    jpeg.Encode(f, res, nil)
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


