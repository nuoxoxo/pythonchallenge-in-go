package main

import (
    "fmt"
    "strings"
    "strconv"
    "reflect"
    "os"
    "image"
    "image/color"
    "image/png"
    _"image/draw"
)

var content string

func main(){

    fmt.Println("init/", content)

    // grap the 2 pts that matter as []str
    first := "first:"
    second := "second:"
    endcomment := "-->"
    fir := strings.Index(content, first)
    sec := strings.Index(content, second)
    end := strings.Index(content, endcomment)
    tmp1 := strings.Split(content[fir + len(first) : sec], ",")
    tmp2 := strings.Split(content[sec + len(second) : end], ",")
    fmt.Println("p1/", tmp1, reflect.TypeOf(tmp1))
    fmt.Println("p2/", tmp2, reflect.TypeOf(tmp2))

    // convert both list to []int - find side len
    p1, p2 := make([]int, len(tmp1)), make([]int, len(tmp2))
    side := -1
    for i, num := range tmp1 {
        n, _ := strconv.Atoi(strings.TrimSpace(num))
        p1[i] = n
        if side < n {
            side = n
        }
    }
    for i, num := range tmp2 {
        n, _ := strconv.Atoi(strings.TrimSpace(num))
        p2[i] = n
        if side < n {
        side = n
        }
    }

    fmt.Println("p1/", p1, reflect.TypeOf(p1))
    fmt.Println("p2/", p2, reflect.TypeOf(p2))
    fmt.Println("side/", side)

    // let's draw
    var i int
    imgBresenham := image.NewRGBA(image.Rect(0, 0, side + 42, side + 42))
    imgNaiveDraw := image.NewRGBA(image.Rect(0, 0, side + 42, side + 42))
    green := color/*.White*/.RGBA{0x80, 0x80, 0x00, 0xff}
    orange := color/*.White*/.RGBA{0xff, 0xa5, 0x00, 0xff}

    i = 0
    for i < len(p1) - 4 {
        Bresenham(imgBresenham, p1[i], p1[i + 1], p1[i + 2], p1[i + 3], green)
        NaiveDraw(imgNaiveDraw, p1[i], p1[i + 1], p1[i + 2], p1[i + 3], green)
        i += 2
    }
    i = 0
    for i < len(p2) - 4 {
        Bresenham(imgBresenham, p2[i], p2[i + 1], p2[i + 2], p2[i + 3], orange)
        NaiveDraw(imgNaiveDraw, p2[i], p2[i + 1], p2[i + 2], p2[i + 3], orange)
        i += 2
    }

    outBresenham, _ := os.Create("output_bresenham.png")
    defer outBresenham.Close()
    png.Encode(outBresenham, imgBresenham)

    outNaiveDraw, _ := os.Create("output_naive.png")
    defer outNaiveDraw.Close()
    png.Encode(outNaiveDraw, imgNaiveDraw)

}


func Bresenham(img *image.RGBA, r, c, rr, cc int, clr color.Color) {
    dr := rr - r
    if dr < 0 {
        dr = -dr
    }
    dc := cc - c
    if dc < 0 {
        dc = -dc
    }
    sr := -1
    sc := -1
    if r < rr {
        sr = 1
    }
    if c < cc {
        sc = 1
    }
    _error := dr - dc
    for {
        img.Set(r, c, clr)
        if r == rr && c == cc {
            break
        }
        err2 := 2 * _error
        if err2 > -dc {
            _error -= dc
            r += sr
        }
        if err2 < dr {
            _error += dr
            c += sc
        }
    }
}

func NaiveDraw(img *image.RGBA, r, c, rr, cc int, clr color.Color) {
    dr := float64(rr - r)
    if dr < 0 {
        dr = -dr
    }
    dc := float64(cc - c)
    if dc < 0 {
        dc = -dc
    }

    if dr > dc { // when line seems more vertical
        if r > rr {
            r, rr = rr, r
            c, cc = cc, c
        }
        m := dc / dr
        y := float64(c)
        for x := r; x <= rr; x++ {
            img.Set(x, int(y), clr)
            y += m
        }
    } else { // or more horizontal
        if c > cc {
            r, rr = rr, r
            c, cc = cc, c
        }
        m := dr / dc
        x := float64(r)
        for y := c; y <= cc; y++ {
            img.Set(int(x), y, clr)
            x += m
        }
    }
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func fabs(a float64) float64 {
    if a < 0 {
        return -a
    }
    return a
}

// use init() to avoid spam

func init(){
    content = `<!--

first+second=?

first:
146,399,163,403,170,393,169,391,166,386,170,381,170,371,170,355,169,346,167,335,170,329,170,320,170,
310,171,301,173,290,178,289,182,287,188,286,190,286,192,291,194,296,195,305,194,307,191,312,190,316,
190,321,192,331,193,338,196,341,197,346,199,352,198,360,197,366,197,373,196,380,197,383,196,387,192,
389,191,392,190,396,189,400,194,401,201,402,208,403,213,402,216,401,219,397,219,393,216,390,215,385,
215,379,213,373,213,365,212,360,210,353,210,347,212,338,213,329,214,319,215,311,215,306,216,296,218,
290,221,283,225,282,233,284,238,287,243,290,250,291,255,294,261,293,265,291,271,291,273,289,278,287,
279,285,281,280,284,278,284,276,287,277,289,283,291,286,294,291,296,295,299,300,301,304,304,320,305,
327,306,332,307,341,306,349,303,354,301,364,301,371,297,375,292,384,291,386,302,393,324,391,333,387,
328,375,329,367,329,353,330,341,331,328,336,319,338,310,341,304,341,285,341,278,343,269,344,262,346,
259,346,251,349,259,349,264,349,273,349,280,349,288,349,295,349,298,354,293,356,286,354,279,352,268,
352,257,351,249,350,234,351,211,352,197,354,185,353,171,351,154,348,147,342,137,339,132,330,122,327,
120,314,116,304,117,293,118,284,118,281,122,275,128,265,129,257,131,244,133,239,134,228,136,221,137,
214,138,209,135,201,132,192,130,184,131,175,129,170,131,159,134,157,134,160,130,170,125,176,114,176,
102,173,103,172,108,171,111,163,115,156,116,149,117,142,116,136,115,129,115,124,115,120,115,115,117,
113,120,109,122,102,122,100,121,95,121,89,115,87,110,82,109,84,118,89,123,93,129,100,130,108,132,110,
133,110,136,107,138,105,140,95,138,86,141,79,149,77,155,81,162,90,165,97,167,99,171,109,171,107,161,
111,156,113,170,115,185,118,208,117,223,121,239,128,251,133,259,136,266,139,276,143,290,148,310,151,
332,155,348,156,353,153,366,149,379,147,394,146,399

second:
156,141,165,135,169,131,176,130,187,134,191,140,191,146,186,150,179,155,175,157,168,157,163,157,159,
157,158,164,159,175,159,181,157,191,154,197,153,205,153,210,152,212,147,215,146,218,143,220,132,220,
125,217,119,209,116,196,115,185,114,172,114,167,112,161,109,165,107,170,99,171,97,167,89,164,81,162,
77,155,81,148,87,140,96,138,105,141,110,136,111,126,113,129,118,117,128,114,137,115,146,114,155,115,
158,121,157,128,156,134,157,136,156,136

-->`
}

