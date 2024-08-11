package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "strconv"
    "reflect"
    "image"
    "image/color"
    "image/gif"
    "image/png"
    "os"
    "bytes"
    "io"
    _"image/draw"
    "math/cmplx"
)

var mandgrey *image.RGBA
var mandblue *gif.GIF
var mdb *image.Gray
var bcolor, gcolor []color.Color

func abs(a, b uint8) uint8 {
    if a > b { return a - b }
    return b - a
}

func main(){
    fmt.Println(Cyan + "len/bcolor" + Rest, len(bcolor),
        Cyan + "len/gcolor" + Rest, len(gcolor),
        Cyan + "len/mdb" + Rest, len(mdb.Pix))
    yell("should be/ 307200")
    N := len(bcolor) // assert len1 == len2
    if N != len(gcolor) { panic("len1 != len2") }
    var r, rr, g, b uint32
    var diff []byte
    for i := 0; i < N; i++ {
        r, g, b, _ = bcolor[i].RGBA()
        r8, g8, b8 := float64(r >> 8), float64(g >> 8), float64(b >> 8)
        grey := uint8(0.299*float64(r8) + 0.587*float64(g8) + 0.114*float64(b8))
        rr, _, _, _ = gcolor[i].RGBA()
        if grey != uint8(rr>>8) {
            diff = append(diff, abs(grey, uint8(rr>>8)))
        }
        //fmt.Println(lum, "-", uint8(rr >> 8))
    }
    fmt.Print("len/", len(diff), diff[:42])
    

/*
    fg, _ := os.Open("mandelbrot.gif")
    defer fg.Close()
    fp, _ := os.Open("mandelbrot_grey.png")
    defer fp.Close()

    dcg, _, _ := image.Decode(fg)
    dcp, _, _ := image.Decode(fp)
    fmt.Println(reflect.TypeOf(dcg), reflect.TypeOf(dcp))

    var buffg, buffp bytes.Buffer
    _ = gif.Encode(& buffg, dcg, nil) // use png.Encode for gif
    //_ = png.Encode(& buffg, dcg) // use png.Encode for gif
    _ = png.Encode(& buffp, dcp)
    fmt.Println(buffg.Bytes()[:42], Yell + "buffer/gif" + Rest, len(buffg.Bytes()))
    fmt.Println(buffp.Bytes()[:42], Yell + "buffer/png" + Rest, len(buffp.Bytes()))
*/
}

func rgb8bit(c color.Color) (byte, byte, byte) {
    r,g,b,_ := c.RGBA()
    return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}

func color2int (src color.Color) uint32 {
    rr,gg,bb, _ := src.RGBA()
    r := uint8(rr >> 8)
    g := uint8(gg >> 8)
    b := uint8(bb >> 8)
    return (uint32(r) << 16) | (uint32(g) << 8) | uint32(b)
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
    var fourfloats [4]float64
    for i, m := range matches[0][1:] {
        val, _ := strconv.ParseFloat(string(m), 64)
        fourfloats[i] = val
    }
    fmt.Println(Cyan + "fourfloats/" + Rest, fourfloats)

    // original mandelbrot.GIF on main page
    prev := sub[:5]
    re = regexp.MustCompile(`(?s)img src="(.*?)"`)
    sub = re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(Cyan + "prev/sub" + Rest, prev, sub)
    data, _ = getbody(prev + sub, "kohsamui", "thailand")
    fmt.Println(string(data)[:123])
    yell("body ends/\n")

    // read as GIF
    reader := bytes.NewReader(data)
    f, _ := os.Create(sub)
    defer f.Close()
    _, _ = io.Copy(f, reader) // save it for later use
    reader.Seek(0, 0)
    mandblue, _ = gif.DecodeAll(reader)
    fmt.Println(Cyan + "original mandblue/typ" + Rest, reflect.TypeOf(mandblue),
        "dim/", mandblue.Image[0].Bounds())
    W, H := mandblue.Image[0].Bounds().Dx(), mandblue.Image[0].Bounds().Dy()
    fmt.Println(Cyan + "original mandblue/w/h" + Rest, W, H)
    // make new mandelbrot
    var L,T,X,Y float64 = fourfloats[0],fourfloats[1],fourfloats[2],fourfloats[3]

    //mandgrey = image.NewGray(image.Rect(0, 0, W, H))
    mandgrey = image.NewRGBA(image.Rect(0, 0, W, H))
    ///fmt.Println(Cyan + "mandgrey/typ" + Rest, reflect.TypeOf(mandgrey))
    mdb = image.NewGray(image.Rect(0, 0, W, H))
    for h := 0; h < H; h++ {
        for w := 0; w < W; w++ {
            bcolor = append(bcolor, mandblue.Image[0].At(w, h))
            realpt := L + float64(w) * X / float64(W)
            imagpt := T + float64(h) * Y / float64(H)
            c := complex( realpt, imagpt )
                //  left + x * width / img.width
                //  top + y * height / img.height
            z := complex(0, 0)
            var i int
            for i = 0; i < 128; i++ {
                z = z * z + c
                if cmplx.Abs(z) > 2 {break}
            }
            grey := uint8(255 * i / 128)
            mandgrey.Set(w, H - 1 - h, color.RGBA{ R: grey, G: grey, B: grey, A: 255 })
            mdb.SetGray(w, h, color.Gray{Y: uint8(255 * i / 128)})
            //gcolor = append(gcolor, color.RGBA{ R: grey, G: grey, B: grey, A: 255 })
            //gcolor = append(gcolor, color.Gray{ Y: grey })
        }
    }
    for h := 0; h < H; h++ {
        for w := 0; w < W; w++ {
            // append here to flatten it
            gcolor = append(gcolor, mandgrey.At(w, h))
        }
    }

    f, _ = os.Create(sub[:len(sub) - 4] + "_grey.png")
    defer f.Close()
    png.Encode(f, mandgrey)

    fmt.Println(L, T, X, Y)
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


