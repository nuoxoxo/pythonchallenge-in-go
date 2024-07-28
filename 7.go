package main

import (
    "fmt"
    "image"
    "image/png"
    _"net/http"
    _"io"
    "os"
    "strings"
    "reflect"
    "strconv"
)

type Pixel struct { R,G,B,A int8 }

func main(){

    /*
    URL := "http://www.pythonchallenge.com/pc/def/oxygen.png"
    resp, _ := http.Get(URL)
    defer resp.Body.Close()

    imgfile, _ := os.Create("oxygen.png")
    defer imgfile.Close()

    _, err := io.Copy(imgfile, resp.Body)
    if err != nil { return }
    */

    imgfile, _ := os.Open("oxygen.png")
    defer imgfile.Close()

    var img image.Image
    img, _ = png.Decode(imgfile)
    bounds := img.Bounds()
    C, R := bounds.Dx(), bounds.Dy()
    fmt.Println("w/", C, "h/", R)

    mid := R / 2
    //set := make(map[Pixel]bool)
    pixels := []Pixel{}
    c := 0
    for c < C {
        color := img.At(c, mid)
        r, g, b, a := color.RGBA() // RGBA is uint32 -> to be converted to 8-bit/0..255
        pixel := Pixel{
            R: int8(r),
            G: int8(g),
            B: int8(b),
            A: int8(a),
        }
        pixels = append(pixels, pixel)
        c += 7
    }
    asc := []int8{}
    for _, p := range pixels {
        if p.R == p.G && p.G == p.B {
            asc = append(asc, p.R)
        }
    }
    //fmt.Println("int/", asc)
    msg := ""
    for _, char := range asc {
        msg += string(char)
    }
    fmt.Println("msg/", msg)

    l := strings.Index(msg, "[")
    r := strings.Index(msg, "]")
    nums := strings.Split(msg[l + 1 : r], ", ")
    fmt.Println("num/", nums, reflect.TypeOf(nums))

    bytes := []byte{}
    for _, num := range nums {
        n, _ := strconv.Atoi(num)
        bytes = append(bytes, byte(n))
    }
    fmt.Println("res/", string(bytes))
}

