package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "strconv"
    "reflect"
    "image"
    _"image/color"
    "image/gif"
    _"image/png"
    _"os"
    "bytes"
    _"io"
    _"image/draw"
    _"math/cmplx"
)

func main(){

}

func init(){
    // main page
    data, _ := getbody("ring/grandpa.html", "repeat", "switch")
    fmt.Println(string(data))
    yell("body ends/\n")

    // href
    re := regexp.MustCompile(`(?s)<a href="../(.*?)"`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(Cyan + "sub/" + Rest, sub)
    data, _ = getbody(sub, "kohsamui", "thailand")
    fmt.Println(string(data))
    yell("body ends/todo\n")
    // get/regex - `<window left="0.34" top="0.57" width="0.036" height="0.027"/>`
    exp := `left="([\d.]+)"\s+top="([\d.]+)"\s+width="([\d.]+)"\s+height="([\d.]+)"`
    re = regexp.MustCompile(exp)
    matches := re.FindAllStringSubmatch(string(data), -1)

    // get the 4 floats
    var fourfloats [4]float64
    for i, m := range matches[0][1:] {
        val, _ := strconv.ParseFloat(string(m), 64)
        fourfloats[i] = val
    }
    fmt.Println(Cyan + "fourfloats/" + Rest, fourfloats)

    // get/original mandelbrot.GIF on main page
    prev := sub[:5]
    re = regexp.MustCompile(`(?s)img src="(.*?)"`)
    sub = re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(Cyan + "prev/sub" + Rest, prev, sub)
    data, _ = getbody(prev + sub, "kohsamui", "thailand")
    fmt.Println(string(data)[:42])
    yell("body ends/\n")

    fmt.Println(string(data)[:42])
    fmt.Println(Yell + "type/data " + Rest, reflect.TypeOf(data))


    bytereader := bytes.NewReader(data)
    img, err := gif.Decode( bytereader )
    if err != nil { fmt.Println("gif.Decode/err", err) }
    imgpal, ok := img.(*image.Paletted)
    if ! ok { fmt.Println("Paletted/not") }
    pal := imgpal.Palette
    bounds := imgpal.Bounds()
    fmt.Println("bounds/", bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            idx := imgpal.ColorIndexAt(x, y)
            col := pal[idx]
            fmt.Println("idx/", idx, "col/", col)
            if x == 42 {return} // to be modif. - TODO
        }
    }

    // TODO - next step is to create the new GIF w/ the given floats and maxIter
}



//

func getbody(sub, u, p string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth(u, p)
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) { fmt.Println( Yell + s + Rest )}
func cyan(s string) { fmt.Println( Cyan + s + Rest )}

