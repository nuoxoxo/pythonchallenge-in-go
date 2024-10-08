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
    "os"
    "bytes"
    "math/cmplx"
)

func main(){

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
    // get/regex
    //  `<window left="0.34" top="0.57" width="0.036" height="0.027"/>`
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

    // read original mandelbrot.GIF on main page
    bytereader := bytes.NewReader(data)
    img, err := gif.Decode( bytereader )
    if err != nil { fmt.Println("gif.Decode/err", err) }
    mandb_paletted, ok := img.(*image.Paletted)
    if ! ok { fmt.Println("Paletted/not") }
    bounds := mandb_paletted.Bounds()
    W, H := bounds.Max.X, bounds.Max.Y
    fmt.Println("bounds/", bounds)

    // PalettedDBG
    PalettedDBG(mandb_paletted, W, H)
    fmt.Println(Yell + "PalettedDBG/ends" + Rest)

    // Step/ make new mandelbrot

    var L,T,X,Y float64 = fourfloats[0],fourfloats[1],fourfloats[2],fourfloats[3]

    // define a greyscale palette with 256 levels of grey
    var greypalette color.Palette
    for i := 0; i < 256; i++ {
        greypalette = append(greypalette, color.Gray{ Y: uint8(i) })
    }
    mandb_newdata := image.NewPaletted(image.Rect(0, 0, W, H), greypalette)

    // Generate the Mandelbrot fractal
    for h := 0; h < H; h++ {
        for w := 0; w < W; w++ {
            realpt := L + float64(w) * (X / float64(W))
            imagpt := T + float64(h) * (Y / float64(H))
            c := complex(realpt, imagpt)
            z := complex(0, 0)
            var i int
            for i = 0; i < 128; i++ {
                z = z*z + c
                if cmplx.Abs(z) > 2 {
                    break
                }
            }
            if i == 128 { i-- }
            var grey uint8 = uint8(i) // BUG - solved
            mandb_newdata.SetColorIndex(w, H - 1 - h, grey)
        }
    }

    outFile, err := os.Create("mandelbrot2.gif")
    if err != nil { panic(err) }
    defer outFile.Close()
    err = gif.Encode(outFile, mandb_newdata, nil)
    if err != nil { panic(err) }

    // PalettedDBG
    PalettedDBG(mandb_newdata, W, H)

    diffs := [][]uint8{}
    for y := 0; y < H; y++ {
        for x := 0; x < W; x++ {
            a := mandb_paletted.ColorIndexAt(x, y)
            b := mandb_newdata.ColorIndexAt(x, y)
            if a != b && b != 128 {
                diffs = append(diffs, []uint8{a, b})
            }
        }
    }

    // DBG
    fmt.Println("len/", len(diffs), "- assert/1679")
    fmt.Println(":21/", diffs[:21])

    N := len(diffs)
    factors := []int{} // finding out factors, assert: should be 2 values
    fac := 2
    for fac < N / 2 + 1 {
        if N % fac == 0 {
            factors = append(factors, fac)
            N /= fac
            break
        }
        fac++
    }
    if len(factors) < 2 {
        factors = append(factors, N)
    }
    fmt.Println("fac/", factors)

    finalW, finalH := factors[0], factors[1]
    finaldata := image.NewPaletted(image.Rect(0, 0, finalW, finalH), greypalette)
    for i, pair := range diffs {
        a, b := pair[0], pair[1]
        pix := 0
        if a > b {
            pix = 234
        }
        finaldata.SetColorIndex(i % finalW, i / finalW, uint8(pix))
    }
    outfile, _ := os.Create("what_.gif")
    defer outfile.Close()
    _ = gif.Encode(outfile, finaldata, nil)
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

func PalettedDBG(data *image.Paletted, W, H int) {
    flag := false
    pal := data.Palette
    for y := 0; y < H; y++ {
        for x := 0; x < W; x++ {
            idx := data.ColorIndexAt(x, H - y - 1)
            col := pal[idx]
            fmt.Println("idx/", idx, "col/", col)
            if x == 42 {flag = true}
        }
        if flag {break}
    }
    fmt.Println(Yell + "Paletted DBG/ends" + Rest)
}

