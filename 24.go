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
    "image/color"
    "os"
    _"log"
    _"archive/zip"
)

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
var img *image.NRGBA

func main(){


    X, Y := img.Bounds().Max.X, img.Bounds().Max.Y
    fmt.Println("typ/png", reflect.TypeOf(img), "bound/", X, Y)
    // start/end
    pair := Blackdots(img)
    fmt.Println("dot/", pair)
    if len(pair) != 2 { panic("wtf/") }
    start, end := pair[1], pair[0]
    ex, ey := end[0], end[1]

    // BFS
    Q := [][2]int{ start }
    SEEN := make(map[[2]int]bool)
    SEEN[ start ] = true
    mapping := make(map[[2]int][2]int)
    dx := []int{-1,0,1, 0}
    dy := []int{ 0,1,0,-1}
    var r, g, b uint32

    for len(Q) > 0 {
        src := Q[0]
        Q = Q[1:]
        fmt.Println(src[0], src[1])
        if src[0] == ex && src[1] == ey {
            break
        }
        for i := 0; i < 4; i++ {
            x, y := src[0] + dx[i], src[1] + dy[i]
            des := [2]int{ x, y }
            if SEEN[des] { continue }
            if !(-1 < x && x < X && -1 < y && y < Y) { continue }
            r, g, b, _ = img.NRGBAAt(x, y).RGBA()
            if r >> 8 != 255 || g >> 8 != 255 || b >> 8 != 255 {
                mapping[des] = src
                SEEN[des] = true
                Q = append(Q, des)
            }
        }
    }
    fmt.Println("s/e", start, end)

    uint32data := []uint32{}
    for end != start {
        r, _, _, _ = img.NRGBAAt(end[0], end[1]).RGBA()
        if r >> 8 != 0 { uint32data = append(uint32data, r) }
        end = mapping[end]
    }

    uint8data := make([]uint8, len(uint32data) * 4)//[]uint8{}
    for i, n := range uint32data {
        uint8data[i * 4] = byte(n >> 24)
        uint8data[i * 4 + 1] = byte(n >> 16)
        uint8data[i * 4 + 2] = byte(n >> 8)
        uint8data[i * 4 + 3] = byte(n)
        /*
        uint8data[i * 4 + 3] = byte(n)
        uint8data[i * 4 + 2] = byte(n >> 8)
        uint8data[i * 4 + 1] = byte(n >> 16)
        uint8data[i * 4] = byte(n >> 24)
        */
    }

    // FIXME
    f, _ := os.Create("out.zip")
    defer f.Close()
    err, _ := f.Write(uint8data)
    fmt.Println(err)
}

func Blackdots(img *image.NRGBA) [][2]int {
    res := [][2]int{}
    X, Y := img.Bounds().Max.X, img.Bounds().Max.Y
    var r, g, b uint32
    for x := 0; x < X; x++ {

        r, g, b, _ = img.NRGBAAt(x, 0).RGBA()
        if r == 0 && g == 0 && b == 0 { res = append(res, [2]int{x, 0}) }

        r, g, b, _ = img.NRGBAAt(x, Y - 1).RGBA()
        if r == 0 && g == 0 && b == 0 { res = append(res, [2]int{x, Y - 1}) }
    }
    return res
}

func init() {

    // get url
    data, _ := getbody("ambiguity.html")
    body := string(data)
    fmt.Println(body)
    yell("body ends/\n")
    // get png
    re := regexp.MustCompile(`(?s)<img src="(.*?)"`)
    sub := re.FindAllStringSubmatch(body, -1)[0][1]
    data, _ = getbody(sub)
    // read png
    reader := bytes.NewReader(data)
    fmt.Println("typ/data", reflect.TypeOf(data), "typ/reader", reflect.TypeOf(reader))
    dec, _ := png.Decode(reader)
    img = dec.(*image.NRGBA)
    //bounds := img.Bounds()
    fmt.Println("typ/png", reflect.TypeOf(img), "bounds/", img.Bounds())

    //Walker(img)

    counter := make(map[color.Color]int)
    for y := 0; y < img.Bounds().Max.Y; y++ {
        for x := 0; x < img.Bounds().Max.X; x++ {
            counter[ img.At(x, y) ]++
        }
    }
    fmt.Println("unique pixel count/", len(counter))
    yell( "init/ends\n" )
}

func yell(s string) { fmt.Println(Yell + s + Rest) }
func cy(s string)   { fmt.Println(Cyan + s + Rest) }

func Walker(img *image.NRGBA) {
    bounds := img.Bounds()
    X, Y := bounds.Max.X, bounds.Max.Y
    for y := 0; y < Y; y++ {
        for x := 0; x < X; x++ {
            r, g, b, _ := img.NRGBAAt(x, y).RGBA()
            fmt.Println("x,y/", x, y, "rgb", r / 257, g / 257, b / 257)
        }
    }
}

func getbody(sub string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/hex/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    return data, nil
}
