package main

import (
    "fmt"
    "image"
    "image/png"
    "net/http"
    "io"
    "os"
    "strings"
    "reflect"
    "strconv"
)

type Pixel struct { R,G,B,A int8 }

func main(){

    PATH := "./oxygen.png"

    defer func(){
        err := os.Remove( PATH )
        fmt.Println("err/del.", err)
    }()

    imgfile, _ := os.Open( PATH)
    defer imgfile.Close()

    var img image.Image
    img, _ = png.Decode(imgfile)
    bounds := img.Bounds()
    C, R := bounds.Dx(), bounds.Dy()
    fmt.Println("w/", C, "h/", R)

    mid := R / 2
    pixels := []Pixel{}
    // finding sizeof a grey segment while making the []Pixel
    lenmap := make(map[int]int)
    var prev Pixel
    runlen, seglen := 0, -1
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
        if c == 0 {
            prev = pixel
        } else {
            if pixel == prev {
                runlen++
            } else {
                lenmap[runlen]++
                runlen = 0
                prev = pixel
            }
        }
        pixels = append(pixels, pixel)
        c++
    }

    maxtimes := -1
    for l, times := range lenmap {
        // fmt.Println(l, ":", times)
        if maxtimes < times {
            maxtimes = times
            seglen = l
        }
    }

    fmt.Println("seg/", seglen)

    asc := []int8{}
    c = 0
    for c < C {
        p := pixels[c]
        if p.R == p.G && p.G == p.B {
            asc = append(asc, p.R)
        }
        c += seglen + 1
    }

    // fmt.Println("int/", asc)

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

func init(){

    URL := "http://www.pythonchallenge.com/pc/def/oxygen.png"
    resp, _ := http.Get(URL)
    defer resp.Body.Close()

    imgfile, _ := os.Create("oxygen.png")
    defer imgfile.Close()

    _, _ = io.Copy(imgfile, resp.Body)

}


