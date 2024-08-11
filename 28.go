package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    _"reflect"
    "regexp"
    "bytes"
    "image/png"
    _"image"
)


func main(){
}

func init(){
    data, _ := getbody("bell.html")
    re := regexp.MustCompile(`(?s)<img src="(.*?)"`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(yell("sub/"), sub)
    //data, _ = getbody("green.html") // x
    //data, _ = getbody("grin.html") // x
    data, _ = getbody(sub) 
    fmt.Println(string(data)[:42], cyan("\nbody ends/"))
    reader := bytes.NewReader(data)
    decoder, err := png.Decode(reader)
    //_, ok := decoder.(*image.Paletted)
    //fmt.Println(cyan("is paletted img?/"), ok)
    fmt.Println(yell("err/png"), err, "len/png", decoder.Bounds())
    //reader.Seek(0, 0)
    X, Y := decoder.Bounds().Max.X, decoder.Bounds().Max.Y
    var G []uint8
    for y := 0; y < Y; y++ {
        for x := 0; x < X; x++ {
            _, g, _, _ := decoder.At(x, y).RGBA()
            G = append(G, uint8(g >> 8))
        }
    }
    var diff []uint8
    var ft []uint8
    for i := 0; i < len(G); i += 2 {
        var dif uint8
        if G[i] < G[i + 1] {
            dif = G[i + 1] - G[i]
        } else {
            dif = G[i] - G[i + 1]
        }
        diff = append(diff, dif)
        //if dif != 42 { ft = append(ft, dif) }
    }
    for _, n := range diff {
        if n != 42 {
            ft = append(ft,n )
        }
    }
    fmt.Println("42/y", string(diff)[100], "len/", len(diff))
    fmt.Println("42/n", string(ft), "len/", len(ft))
}

func getbody(sub string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/ring/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "repeat", "switch" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    if len(data) < 123 { fmt.Println(string(data), yell("\bbody ends/")) }
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) string { return Yell + s + Rest }
func cyan(s string) string { return Cyan + s + Rest }


