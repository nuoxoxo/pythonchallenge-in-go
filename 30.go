package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
//)/*
    "image"
    "image/color"
    "image/png"
    "os"
    "strconv"
)//*/

func main(){
    data, _ := getbody("yankeedoodle.html")
    data, _ = getbody("yankeedoodle.csv")
    page := strings.ReplaceAll(string(data), "\n", "")
    var fwords []string
    for _, line := range strings.Split(page, ",") {
        fwords = append(fwords, strings.TrimSpace(line))
    }
    fmt.Println(yell("fwords/"), fwords[:42])
    fmt.Println(yell("len/"), len(fwords))
    var factors []int
    for n := len(fwords) - 1; n > 1; n-- {
        if len(fwords) % n == 0 {
            factors = append(factors, n)
        }
    }
    if len(factors) != 2 { panic("!=2") }
    fmt.Println(yell("fact/"), factors)

    // conversion: []float32
    floats := make([]float32, len(fwords))
    for i, word := range fwords {
        f, _ := strconv.ParseFloat(word, 32)
        floats[i] = float32(f)
    }

    // write to image
    outfile, _ := os.Create("lv30.png")
    defer outfile.Close()
    imgout := image.NewGray(image.Rect(0, 0, factors[1], factors[0]))
    for i, f := range floats {
        r := i / factors[1]
        c := i % factors[1]
        grey := uint8(f * 255)
        imgout.SetGray(c, r, color.Gray{Y: grey})
    }
    png.Encode(outfile, imgout)

    // reversed image/// do it again
    outfile, _ = os.Create("lv30r.png")
    defer outfile.Close()
    imgout = image.NewGray(image.Rect(0, 0, factors[0], factors[1]))
    for i, f := range floats {
        r := i / factors[1]
        c := i % factors[1]
        grey := uint8(f * 255)
        imgout.SetGray(r, c, color.Gray{Y: grey})
    }
    png.Encode(outfile, imgout)

    // n = str(x[i])[5] + str(x[i+1])[5] + str(x[i+2][]6)
    var res []uint8
    for i := 0; i < len(fwords) - 3; i += 3 {
        s := string(fwords[i][5]) + string(fwords[i + 1][5]) + string(fwords[i + 2][6])
        n, _ := strconv.Atoi(s)
        if n > 255 { panic(">255") }
        res = append(res, uint8(n))
    }
    fmt.Println("res/:42", res[:42], len(res))
    fmt.Println("res/str", string(res))

}

func getbody(sub string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/ring/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "repeat", "switch" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(data), yell("\bbody ends/"))
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) string { return Yell + s + Rest }
func cyan(s string) string { return Cyan + s + Rest }

