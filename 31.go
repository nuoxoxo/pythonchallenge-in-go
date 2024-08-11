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
    "image/draw"
    "math/cmplx"
)

var mandgrey *image.RGBA
var mandblue *gif.GIF

func main(){
    fmt.Println("bounds/", mandgrey.Bounds(), mandblue.Image[0].Bounds()) 

    bounds := mandgrey.Bounds()
    Y, X := bounds.Dy(), bounds.Dx()

    //X, Y := mandgrey.Bounds().Dx(), mandgrey.Bounds().Dy()
    mb := image.NewRGBA( mandblue.Image[0].Bounds() )
    mg := image.NewRGBA( mandgrey.Bounds() ) // now both are *image.RGBA

    fmt.Println(Cyan+"mb/Image type/"+Rest, reflect.TypeOf(mandblue.Image))
    fmt.Println(Cyan+"mb/Image type[0]/"+Rest, reflect.TypeOf(mandblue.Image[0]))
    draw.Draw(mb, bounds, mandblue.Image[0], bounds.Min, draw.Src)
    draw.Draw(mg, bounds, mandgrey, bounds.Min, draw.Src)
    /*var diff []struct {
        X, Y  int
        Color1, Color2 color.Color
	}*/
    var diff [][2]uint8
    diffmap := make(map[[2]uint8]int)
    var r,g,b uint32
    for y := 0; y < Y; y++ {
        for x := 0; x < X; x++ {
            r, g, b, _ = mb.At(x, y).RGBA()
            /*
            L := uint8(0.299 * float64(r >> 8)) +
                uint8(0.587 * float64(g >> 8)) +
                uint8(0.114 * float64(b >> 8))
            */
            L := uint8(0.2126 * float64(r >> 8)) +
                uint8(0.7152 * float64(g >> 8)) +
                uint8(0.0722 * float64(b >> 8))
            r, _, _, _ = mg.At(x, y).RGBA()
            R := uint8(r >> 8)
            //fmt.Println(mb.At(x,y), mg.At(x,y))//.RGBA())
            if L != R {
                //fmt.Println(L, Yell+"vs/"+Rest, R, L - R)
                diff = append(diff, [2]uint8{L, R})
                diffmap[[2]uint8{L, R}]++
            }
        }
    }
    fmt.Println(diff[:42])
    fmt.Println(len(diff), "/should be around 1600")
    for k, v := range diffmap {
        fmt.Println("diff/", k, v)
    }
    fmt.Println("len/", len(diffmap))
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
    for h := H - 1; h > -1; h-- {
    //for h := 0; h < H; h++ {
        for w := 0; w < W; w++ {
            realpt := L + float64(w) * X / float64(W)
            imagpt := T + float64(h) * Y / float64(H)
            c := complex( realpt, imagpt )
            z := complex(0, 0)
            var i int
            for i = 0; i < 128; i++ {
                z = z * z + c
                if cmplx.Abs(z) > 4 {break}
                //if real(z) * real(z) + imag(z) * imag(z) > 4 {break}
                grey := uint8(255 * i / 128)
                mandgrey.Set(w, H - 1 - h, color.RGBA{ R: grey, G: grey, B: grey, A: grey })
            }
        }
    }
    f, _ = os.Create(sub[:len(sub) - 4] + "_grey.png")
    defer f.Close()
    png.Encode(f, mandgrey)
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


