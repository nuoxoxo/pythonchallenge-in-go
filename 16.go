package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "image"
    "image/gif"
    "os"
    "reflect"
)

func main(){
    mozart, _ := os.Open("files/mozart.gif")
    defer mozart.Close()
    mozartgif, _ := gif.Decode(mozart)
    bounds := mozartgif.Bounds()
    fmt.Println("bounds/", bounds, reflect.TypeOf(bounds))
    fmt.Println("rows", bounds.Dy())
    fmt.Println("cols", bounds.Dx())
    R := bounds.Dy()
    C := bounds.Dx()

    // paletted := image.NewPaletted(bounds, mozartgif.(*image.Paletted).Palette)
    pimg, _ := mozartgif.(*image.Paletted)
    r := 0
    for r < R {
        row := make(map[uint8]int)
        c := 0
        for c < C {
            color := pimg.ColorIndexAt(c, r)
            row[color]++
            c++
        }
        for k, v := range row {
            fmt.Println("row/", k, v)
        }
        fmt.Println("")
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

