package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "reflect"
    "bytes"
    "image"
    "image/png"
)

const URL string = "http://www.pythonchallenge.com/pc/hex/"
const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"

func main(){}

func init() {

    // get url
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + "ambiguity.html", nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    body := string(data)
    fmt.Println(body, "\nbody ends/")
    // get png
    re := regexp.MustCompile(`(?s)<img src="(.*?)"`)
    sub := re.FindAllStringSubmatch(body, -1)[0][1]
    conn = & http.Client{}
    req, _ = http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ = conn.Do(req)
    defer resp.Body.Close()
    // read png
    data, _ = ioutil.ReadAll(resp.Body)
    reader := bytes.NewReader(data)
    fmt.Println("typ/data", reflect.TypeOf(data), "typ/reader", reflect.TypeOf(reader))
    img, _ := png.Decode(reader)
    bounds := img.Bounds()
    fmt.Println("typ/png", reflect.TypeOf(img), "bounds/", bounds)
    //if tmp, ok := img.(*image.NRGBA); ok { inspect (tmp) }
}

func inspect(img *image.NRGBA) {
    bounds := img.Bounds()
    X, Y := bounds.Max.X, bounds.Max.Y
    var x, y int
    y = 0
    for y < Y {
        x = 0
        for x < X {
            r, g, b, _ := img.At(x, y).RGBA()
            fmt.Println("x,y/", x, y, "rgb", r / 257, g / 257, b / 257)
            x++
        }
        y++
    }
}



