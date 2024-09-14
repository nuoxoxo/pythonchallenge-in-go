package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "reflect"
    "bytes"
    "image/jpeg"
    "image/png"
    "image/color"
    "image"
    "math"
    "os"
    "strconv"
)

var beer1, beer2 []uint8

func main() {

    // work with beer2, some steps are duplicated from init()
    beer2_reader := bytes.NewReader(beer2)
    beer2_decoder, _ := png.Decode( beer2_reader )
    bounds := beer2_decoder.Bounds()
    W, H := bounds.Max.X, bounds.Max.Y
    file_index := 0

    // 1st run?
    pixels := []uint32{}
    for y := 0; y < H; y++ {
        for x := 0; x < W; x++ {
            color := beer2_decoder.At(x, y)
            scale, _, _, _ := color.RGBA()
            scale /= 257 // uint32
            pixels = append(pixels, scale)
        }
    }

    for true {
        // get the brightest pixel
        max := uint32(0)
        for _, pixel := range pixels {
            if max < pixel {
                max = pixel
            }
        }
        // re-group pixels xcpt the brightest one
        temp := []uint32{}
        for _, pixel := range pixels {
            if pixel < max {
                temp = append(temp, pixel)
            }
        }
        pixels = temp
        if len(pixels) == 0 {
            break
        }
        fmt.Println("\nmax/", max)
        fmt.Println("ttl/", len(pixels))
        sqrt := math.Sqrt(float64(len(pixels)))
        side := int(sqrt)
        fmt.Println("sqrt/", sqrt)
        fmt.Println("side/", side)

        // draw a res png
        new_image := image.NewRGBA(image.Rect(0, 0, side, side))
        pixel_index := 0
        for y := 0; y < side; y++ {
            for x := 0; x < side; x++ {
                scale := uint8(pixels[pixel_index])
                //scale := pixels[pixel_index]
                color2set := color.RGBA{ scale, scale, scale, 255 }
                new_image.Set(x, y, color2set)
                pixel_index++
            }
        }
        file, _ := os.Create(strconv.Itoa(file_index) + ".png")
        defer file.Close()
        _ = png.Encode(file, new_image)
        file_index++
    }
}

func init() {

    data, _ := getbody("rock/beer.html", "kohsamui", "thailand")
    fmt.Println(string(data))

    // GET beer1.jpg
    re := regexp.MustCompile(`(?s)<img src="(.*?)"`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    beer1, _ = getbody("rock/" + sub, "kohsamui", "thailand")
    fmt.Println(Cyan + "\nsub/data on newline" + Rest)
    fmt.Println(string( beer1 [:123]))
    fmt.Println(reflect.TypeOf( beer1 ), sub)

    beer1_reader := bytes.NewReader(beer1)
    beer1_decoder, err := jpeg.Decode( beer1_reader )
    if err != nil {
        fmt.Println("jpeg.Decode/err", err)
    }
    bounds := beer1_decoder.Bounds()
    W, H := bounds.Max.X, bounds.Max.Y
    fmt.Println("bounds/", bounds, "WH", W, H)

    // GET beer2.png
    beer2, _ = getbody("rock/" + "beer2.png", "kohsamui", "thailand")
    fmt.Println(Cyan + "\nsub/data on newline" + Rest)
    fmt.Println(string( beer2 [:123]))
    fmt.Println(reflect.TypeOf( beer2 ))

    beer2_reader := bytes.NewReader(beer2)
    beer2_decoder, err := png.Decode( beer2_reader )
    if err != nil {
        fmt.Println("png.Decode/err", err)
    }
    bounds = beer2_decoder.Bounds()
    W, H = bounds.Max.X, bounds.Max.Y
    fmt.Println("bounds/", bounds, "WH", W, H)

    // check pixels
    dict := make(map[uint32]int)
    for y := 0; y < H; y++ {
        for x := 0; x < W; x++ {
            color := beer2_decoder.At(x, y)
            r, g, b, _ := color.RGBA()
            // assert greyscale
            if !(r == g && g == b) {
                panic("checking/greyscale")
            }
            fmt.Println(x, y, "-", r / 257)
            dict [r / 257]++
        }
    }
    fmt.Println(dict, "/dict")

    fmt.Println(Cyan + "init/ends \n" + Rest)
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



