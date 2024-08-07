package main

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    "regexp"
    "strings"
    "encoding/base64"
    "reflect"
)

var base64string string
var URL, BODY, Filename string
const offset int = 64

func main(){

    // now get the .wav sound file
    base64uint8, _ := base64.StdEncoding.DecodeString( base64string )
    fmt.Println("base64 data/head:", base64uint8[ :offset], reflect.TypeOf(base64uint8))

    // big--->little endian 
    lendian := []uint8{}
    idx := 44
    N := len(base64uint8)
    wave := base64uint8[ :idx] // idea/ cp the wav-header not gonna reverse it
    lendian = append(lendian, wave...)
    for {
        end := idx + 4
        if end > N {
            end = N
        }
        wave = base64uint8[idx : end]
        if len(wave) < 4 {
            if len(wave) != 0 { panic("wtf/") }
            break
        }
        i := 3
        for i > -1 {
            lendian = append(lendian, wave[i])
            i--
        }
        idx += 4
    }

    _ = os.WriteFile( Filename , base64uint8, 0644) // big
    _ = os.WriteFile( "endian.wav" , lendian, 0644) // little

}

func init(){

    URL = "http://www.pythonchallenge.com/pc/hex/bin.html"
    ups, sep := "butterfly", 6
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(ups[: sep], ups[sep :])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    BODY = string(body)

    // fmt.Println(BODY, "\nbody ends/\n\n") // very long entire body

    // get filename
    re := regexp.MustCompile(`(?s)name="(.*?)"`)
    matches := re.FindAllStringSubmatch(BODY, -1)
    Filename = matches[0][1]
    fmt.Println("Filename/", Filename)

    // get the bound
    re = regexp.MustCompile(`(?s)boundary="(.*?)"`)
    matches = re.FindAllStringSubmatch(BODY, -1)
    bound := "--" + matches[0][1]
    N := len(matches[0])
    fmt.Println("len/", len(matches[0]), "matches/", matches)
    fmt.Println("match/", strconv.Quote( matches[0][1] ))
    fmt.Println("modf./", strconv.Quote( bound ))
    fmt.Println("check/", "--===============1295515792==")

    // get bounded trunk which should look base64 encoded
    re = regexp.MustCompile(fmt.Sprintf(`(?s)%s(.*?)%s`, bound, bound))
    matches = re.FindAllStringSubmatch(BODY, -1)
    N = len(matches[0][1])
    end := N - offset
    fmt.Println("\nlen/", len(matches[0]))
    fmt.Println("aft/", matches[0][1][: offset], "bef/", matches[0][1][end :])

    trunk := strings.Split(matches[0][1], "\n\n")
    base64string = trunk[1]
    N = len( base64string )
    end = N - offset
    fmt.Println("len/", len(trunk))
    fmt.Println("aft/", trunk[1][: offset], "bef/", trunk[1][end :])
}

