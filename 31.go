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
)



func main(){
    fg, _ := os.Open("mandelbrot.gif")
    defer fg.Close()
    fp, _ := os.Open("mandelbrot_grey.png")
    defer fp.Close()

    dcg, _, _ := image.Decode(fg)
    dcp, _, _ := image.Decode(fp)
    fmt.Println(reflect.TypeOf(dcg), reflect.TypeOf(dcp))

    var buffg, buffp bytes.Buffer
    _ = png.Encode(& buffp, dcp)
    _ = png.Encode(& buffg, dcg) // use png.Encode for gif
    fmt.Println(buffg.Bytes()[:42], Yell + "buffer/gif" + Rest, len(buffg.Bytes()))
    fmt.Println(buffp.Bytes()[:42], Yell + "buffer/png" + Rest, len(buffp.Bytes()))

    //var diff [][2]uint8
    
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

    // original mandelbrot.gif on main page
    prev := sub[:5]
    re = regexp.MustCompile(`(?s)img src="(.*?)"`)
    sub = re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(Cyan + "prev/sub" + Rest, prev, sub)
    data, _ = getbody(prev + sub, "kohsamui", "thailand")
    fmt.Println(string(data)[:123])
    yell("body ends/\n")
    f, _ := os.Create(sub)
    defer f.Close()

    // read as GIF
    reader := bytes.NewReader(data)
    _, _ = io.Copy(f, reader) // save it for later use
    reader.Seek(0, 0)
    mandel, _ := gif.DecodeAll(reader)
    fmt.Println("mandel/typ", reflect.TypeOf(mandel),
        "dim/", mandel.Image[0].Bounds())
    W, H := mandel.Image[0].Bounds().Dx(), mandel.Image[0].Bounds().Dy()
    fmt.Println("w/h", W, H)

    // make new mandelbrot
    var L,T,X,Y float64 = fourfloats[0],fourfloats[1],fourfloats[2],fourfloats[3]
    mdb := image.NewGray(image.Rect(0, 0, W, H))
    for w := 0; w < W; w++ {
        for h := 0; h < H; h++ {
            realpt := L + float64(w) * X / float64(W)
            imagpt := T + float64(h) * Y / float64(H)
            c := complex( realpt, imagpt )
            z := complex(0, 0)
            var i int
            for i = 0; i < 128; i++ {
                z = z * z + c
                if real(z) * real(z) + imag(z) * imag(z) > 4 {break}
                mdb.SetGray(w, h, color.Gray{Y: uint8(255 * i / 128)})
            }
        }
    }
    f, _ = os.Create(sub[:len(sub) - 4] + "_grey.png")
    defer f.Close()
    png.Encode(f, mdb)
}

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

