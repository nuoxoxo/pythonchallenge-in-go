package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "image"
    "image/gif"
    "image/color"
    "os"
    "reflect"
    "strconv"
)

func main(){
    gf, _ := os.Open("files/mozart.gif")
    defer gf.Close()
    mozart, _ := gif.Decode(gf)
    bounds := mozart.Bounds()
    var R, C, r, c int
    R, C = bounds.Dy(), bounds.Dx()
    fmt.Println("size/", bounds, reflect.TypeOf(bounds))
    fmt.Println("rows/", R, "cols/", C)
    res := image.NewPaletted( bounds, nil )
    fmt.Println("init/", /*res,*/ reflect.TypeOf(res))
    
    r = 0
    for r < R {
        // row reprented as a rgb(a) slice
        row := make([]color.Color, C)
        c = 0
        for c < C {
            cl := mozart.At(c, r)
            //fmt.Println("color/", cl, reflect.TypeOf(cl))
            row[c] = cl
            c++
        }
        // here we do longest uni-char substring 
        // todo... doing...
        var commoncolor color.Color = row[0] // most seen
        var currentcolor color.Color = commoncolor // now-inspecting
        scurr := 0 // curr s(tart)
        s, e, maxlen := 0, 0, 0
        c = 1
        for c < C {
            if currentcolor != row[c] {
                dist := c - scurr
                if maxlen < dist {
                    commoncolor = row[c]
                    s = scurr
                    e = c
                    maxlen = dist
                }
                currentcolor = row[c]
                scurr = c
            }
            c++
        }
        se := strconv.Itoa(s) + "-" + strconv.Itoa(e)
        fmt.Println(r, "\b/", "color/", commoncolor, se, "len/", maxlen)
        r++
        // obsevation/
        //  - lots of len-5 segments, not seem to be pink
        //  - some rows dont have any len-5 segments, surprise
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

