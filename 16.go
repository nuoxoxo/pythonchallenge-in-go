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
const offset int = 5
const offl int = 0
const offr int = 0

func main(){
    gf, _ := os.Open("files/mozart.gif")
    defer gf.Close()
    mozart, _ := gif.Decode(gf)
    bounds := mozart.Bounds()
    var R, C, r, c int
    R, C = bounds.Dy(), bounds.Dx()
    fmt.Println("size/", bounds, reflect.TypeOf(bounds))
    fmt.Println("rows/", R, "cols/", C)

    // try something like Paletted
    palettedImg, _ := mozart.(*image.Paletted)

    res := image.NewPaletted( bounds, palettedImg.Palette )
    // res := image.NewPaletted( bounds, nil )
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

        // 1 - finding out about that strip|segment
        //  observation/
        //  {255 0 255 255} --> pink-ish, len-5 segment found
        ///*s, e :=*/findingLongestSegment(row, C)

        // 2 - find index-pairs of all pink segments
        //  assert/
        //  assert/there is only one such pair
        //  once we see a |255,0,255,255| we do e = s + 5 and i += 5
        s := findingPinkSegment(row, C)

        // 3 - move pink segments at the end of row (ie. res[r])
        row = append(row[s:], row[:s]...)
        // row = append(append(row[s:e], row[:s]...), row[e:]...)

        c = 0
        for c < C {
            res.Set(c, r, row[c])
            c++
        }
        r++
        // . . . . . .
    }
    outfile, _ := os.Create("res.gif")
    defer outfile.Close()
    _ = gif.Encode(outfile, res, nil)
    // done?
}


func findingPinkSegment(row[]color.Color, C int)(int/*,int*/)/*(color.Color,int,int,int,string)*/{

    s, e := 0, 0
    c := 1
    for c < C {
        cl := row[c]
        rr, gg, bb, aa := cl.RGBA()
        r8, g8, b8, a8 := uint8(rr>>8), uint8(gg>>8), uint8(bb>>8), uint8(aa>>8)
        // assert curr==next is sufficient to target a 5-pix segment
        if r8 == 255 && g8 == 0 && b8 == 255 && a8 == 255 && row[c]==row[c+1] && row[c+1]==row[c+2] {
            // fmt.Println(cl)
            //return s
            s = c
            e = c + 5
            c += 5
        } else {
            c++
        }
        //c++
    }
    fmt.Println(row[s], s, e, e - s)
    return s//, e
}


func findingLongestSegment(row[]color.Color, C int)(int,int)/*(color.Color,int,int,int,string)*/{

    // here we do longest uni-char substring 
    var commoncolor color.Color = row[0] // most seen
    var currentcolor color.Color = commoncolor // now-inspecting
    scurr := 0 // curr s(tart)
    s, e, maxlen := 0, 0, 0
    c := 1
    for c < C {
        if currentcolor != row[c] {
            dist := c - scurr
            if maxlen < dist {
                // crucial BugFix
                commoncolor = currentcolor// row[c] // BUG
                s = scurr
                e = c
                maxlen = dist
            }
            currentcolor = row[c]
            scurr = c
        }
        c++
    }
    SE := strconv.Itoa(s) + "-" + strconv.Itoa(e)
    //fmt.Println(r, "\b/", "color/", commoncolor, SE, "len/", maxlen)
    fmt.Println(commoncolor, SE, "len/", maxlen)
    return s, e
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

