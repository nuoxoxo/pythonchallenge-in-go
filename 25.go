package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    _"reflect"
    "image"
    "image/draw"
    "image/color"
    "image/png"
    "os"
)

func main(){
    N := getlastindex(1, ".wav")
    waves := getwaves(N, ".wav")
    
    // flattened/ one long continous data w/o wave header
    idx := 44
    flattened := []uint8{}
    for _, wave := range waves {
        data := wave[idx:]
        flattened = append(flattened, data...)
    }
    fmt.Println(len(flattened), Cyan + "should be/ 270000" + Rest)

    // idea/
    //  total bytes: 270000 - 1 pixel = 3 bytes (r,g,b)
    //      ---> 270000 / 3 = 90000 pixels
    //  this is level 25 and there are 25 pcs on lake.jpg
    //      ---> 90000 / 25 = 3600 pixels/image
    //      ---> deduction/each image is 60 x 60
    // todo/
    //  working on a flattened slice/
    //      given a index out of 270000, f(i) = x, y

    // way1
    working_on_nested_slice(waves, "_nested")
    // way2
    working_on_flattened(flattened, "_flattened")
}

// 2 ways

func working_on_flattened(W []uint8, filename string) {

    bytesPerSquarePiece := 10800 // square piece 60 * 60 * rgb 3

    outimage := image.NewRGBA(image.Rect(0, 0, 60 * 5, 60 * 5))
    for i := 0; i < 25; i++ {
        indexstart := bytesPerSquarePiece * i
        img60x60 := image.NewRGBA(image.Rect(0, 0, 60, 60))
        for y := 0; y < 60; y++ {
            for x := 0; x < 60; x++ {
                index := indexstart + (60 * y + x) * 3
                R, G, B := W [index], W [index + 1], W [index + 2]
                img60x60.Set(x, y, color.RGBA{R:R, G:G, B:B, A:255})
            }
        }
        x := 60 * (i % 5)
        y := 60 * (i / 5)
        draw.Draw(outimage, image.Rect(x, y, x + 60, y + 60),
            img60x60, image.Point{0, 0}, draw.Over)
    }
    outfile, _ := os.Create(filename + ".png")
    defer outfile.Close()
    png.Encode(outfile, outimage)    
}

func working_on_nested_slice(waves [][]uint8, filename string) {

    outimage := image.NewRGBA(image.Rect(0, 0, 60 * 5, 60 * 5))
    for i, wave := range waves {
        data := wave[ 44: ]
        img60x60 := image.NewRGBA(image.Rect(0, 0, 60, 60))
        for y := 0; y < 60; y++ {
            for x := 0; x < 60; x++ {
                index := (60 * y + x) * 3
                R, G, B := data[index], data[index + 1], data[index + 2]
                img60x60.Set(x, y, color.RGBA{R:R, G:G, B:B, A:255})
            }
        }
        x := 60 * (i % 5)
        y := 60 * (i / 5)
        draw.Draw(outimage, image.Rect(x, y, x + 60, y + 60),
            img60x60, image.Point{0, 0}, draw.Over)
    }
    outfile, _ := os.Create(filename + ".png")
    defer outfile.Close()
    png.Encode(outfile, outimage)    
}

// helper

func getwaves(N int, sub string) [][]uint8 {
    waves := make([][]byte, N)
    for i := 1; i < N + 1; i++ {
        conn := & http.Client{}
        req, _ := http.NewRequest("GET", URL + strconv.Itoa(i) + sub, nil)
        req.SetBasicAuth("butter", "fly")
        resp, _ := conn.Do(req)
        defer resp.Body.Close()
        data, _ := ioutil.ReadAll(resp.Body)
        waves[i - 1] = data
    }
    return waves
}

func getlastindex(i int, sub string) int {
    for ;;i++ {
        conn := & http.Client{}
        req, _ := http.NewRequest("GET", URL + strconv.Itoa(i) + sub, nil)
        req.SetBasicAuth("butter", "fly")
        resp, _ := conn.Do(req)
        defer resp.Body.Close()
        if resp.StatusCode != 200 {
            fmt.Println("status/ ", resp.StatusCode)
            break
        }
        //data, _ := ioutil.ReadAll(resp.Body)
        //fmt.Println("len/typ", len(data), reflect.TypeOf(data), i)
    }
    return i - 1
}

const URL string = "http://www.pythonchallenge.com/pc/hex/lake"//1.jpg"
const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) { fmt.Println(Yell + s + Rest) }
func cy(s string)   { fmt.Println(Cyan + s + Rest) }


