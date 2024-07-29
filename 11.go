package main

import (
    _"fmt"
    "os"
    "image"
    "image/jpeg"
)

func main() {

    imgfile, _ := os.Open("files/cave.jpg")
    defer imgfile.Close()

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
    imgfileEven, _ := os.Create("files/1.jpg")
    defer imgfileEven.Close()
    jpeg.Encode( imgfileEven, Even, nil )

    imgfileOdd, _ := os.Create("files/2.jpg")
    defer imgfileOdd.Close()
    jpeg.Encode (imgfileOdd, Odd, nil)
}

