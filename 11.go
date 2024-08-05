package main

import (
    "fmt"
    "os"
    "io"
    "net/http"
    "image"
    "image/jpeg"
    _"io/ioutil"
)

const Filename string = "cave.jpg"

func main() {

    imgfile, _ := os.Open( Filename )
    defer imgfile.Close()

    defer func(){
        err := os.Remove(Filename)
        fmt.Println("err/del.", err)
    }()

    img, _, _ := image.Decode( imgfile )
    bounds := img.Bounds()
    Even := image.NewRGBA( bounds )
    Odd := image.NewRGBA( bounds )
    W := bounds.Dx()
    H := bounds.Dy()
    var r, c int
    r = 0
    for r < H {
        c = 0
        for c < W {
            pix := img.At(c, r)
            rc := r + c
            if rc % 2 == 0 {
                Even.Set( c, r, pix )
            } else {
                Odd.Set( c, r, pix )
            }
            c++
        }
        r++
    }

    // save both jpg
    imgfileEven, _ := os.Create("1.jpg")
    defer imgfileEven.Close()
    jpeg.Encode( imgfileEven, Even, nil )

    imgfileOdd, _ := os.Create("2.jpg")
    defer imgfileOdd.Close()
    jpeg.Encode (imgfileOdd, Odd, nil)
}

func init(){

    f, err := os.Create( Filename )
    fmt.Println("err/init", err)
    defer f.Close()

    URL := "http://www.pythonchallenge.com/pc/return/"
    ups, mid := "hugefile", 4

    conn := &http.Client{}
    req, err := http.NewRequest("GET", URL + Filename, nil)
    fmt.Println("err/GET", err)

    req.SetBasicAuth(ups[:mid], ups[mid:])
    resp, err := conn.Do( req )
    fmt.Println("err/GET", err, resp.StatusCode)
    defer resp.Body.Close()

    _, _ = io.Copy(f, resp.Body)
}

