package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "reflect"
    "bytes"
    "image"
    "image/gif"
    "image/color"
)

var rawgif []byte
const URL string = "http://www.pythonchallenge.com/pc/hex/"
const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"

func main(){

    fmt.Println("typ/", reflect.TypeOf(rawgif), "len/", len(rawgif))

    // treat/read data as gif
    reader := bytes.NewReader(rawgif)
    readgif, _ := gif.DecodeAll(reader)
    fmt.Println("typ/reader", reflect.TypeOf(reader))
    fmt.Println("typ/readgif", reflect.TypeOf(readgif))

    // count frames, all frames have same bounds
    nframes := 0
    frameset := make(map[image.Rectangle]bool)

    pxlmap := make(map[color.Color]int)
    ncenter := 0
    xoffset, yoffset, offset := 1000, 1000, 98 // now we know 98,98
    xbound, ybound := -1, -1

    // we also know 100,100 happens 5 times
    FiveGuys := [][][]int{}
    for i := 0; i < 5; i++ { FiveGuys = append(FiveGuys, [][]int{}) }

    for _, frame := range readgif.Image {
        nframes++
        frameset[frame.Bounds()] = true

        // inspect all pixels
        X, Y := frame.Bounds().Max.X, frame.Bounds().Max.Y
        x, y := 0, 0
        for y < Y {
            x = 0
            for x < X {
                r, g, b, _ := frame.At(x, y).RGBA()
                if r / 257 == 8 && g / 257 == 8 && b / 257 == 8 {
                    fmt.Println(Cyan + "color 8/" + Rest, x, y)
                    if x == 100 && y == 100 { ncenter++ }
                    if xoffset > x { xoffset = x } // know a priori:
                    if yoffset > y { yoffset = y } //   offset/ 98,98
                    if xbound < x { xbound = x }
                    if ybound < y { ybound = y }

                    FiveGuys[ncenter-1] = append(
                        FiveGuys[ncenter-1], []int{x - offset, y - offset})
                }
                pxlmap[frame.At(x, y)]++
                x++
            }
            y++
        }
    }
    fmt.Println("nframes/", nframes, "set/", frameset)
    fmt.Println(Yell + "\t^ a total of 133 frames" + Rest)
    for k, v := range pxlmap { fmt.Println("color/", k, "qty", v) }
    fmt.Println(Yell + "\t^ one of 2 existing colors has but 1-pix per frame" + Rest)
    fmt.Println("ncenter/", "coor(100, 100) reached", ncenter, "times")
    fmt.Println("offsets/", xoffset, yoffset)
    fmt.Println("rbounds/", xbound, ybound)

    // Bruteforce
    xbound, ybound = xbound - xoffset + 1, ybound - yoffset + 1
    for _, coors := range FiveGuys {
        char := [][]string{}
        r := 0
        for r < ybound {
            temp := []string{}
            c := 0
            for c < xbound {
                temp = append(temp, " ")
                c++
            }
            r++
            char = append(char, temp)
        }
        for _, coor := range coors {
            x, y := coor[0], coor[1]
            char[x][y] = "@"
        }
        r = 0
        for r < ybound {
            c := 0
            for c < xbound {
                fmt.Print(char[r][c])
                c++
            }
            fmt.Println()
            r++
        }
        fmt.Println("---")
    }
}

func init(){

    // GET
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + "copper.html", nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    // to be told how url should be modified
    temp, _ := ioutil.ReadAll(resp.Body)
    body := string(temp)

    // get sub2, ie. white.gif
    re := regexp.MustCompile(`(?s)maybe (.*?) would`)
    matches := re.FindAllStringSubmatch(body, -1)
    sub2 := matches[0][1]

    conn = & http.Client{}
    req, _ = http.NewRequest("GET", URL + sub2, nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ = conn.Do(req)
    defer resp.Body.Close()

    rawgif, _ = ioutil.ReadAll(resp.Body)
}


