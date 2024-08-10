package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    _"reflect"
    "regexp"
    "bytes"
    "image/png"
    "image/gif"
    "image"
    "image/color"
    "image/draw"
    "reflect"
    "os"
    "compress/bzip2"
    "strings"
    "os/exec"
)

var data []uint8
var pimg *image.Paletted

func main(){
    if pimg == nil { panic("wtf/") }
    fmt.Println(yell("gif/typ"), reflect.TypeOf(pimg), pimg.Bounds())
    var bloc string
    palette := [256]uint8{}
    //palette := make(map[int]uint8)
    for i, pix := range pimg.Palette {
        r,g,b := rgb8bit(pix)
        if r != g || g != b { panic("assert/all pix are grey") }
        // observation/greyscale means only one channel of r,g,b matters 
        if i % 2 == 0 { bloc += Cyan } else { bloc += Yell }
        bloc += fmt.Sprintf("%d %d %d ", r, g, b)
        palette[i] = g
    }
    fmt.Println(bloc, Rest)
    diff := []int{}
    flattened := pimg.Pix
	lookup := palettelookup(flattened, palette)
    for i := 1; i < len(flattened); i++ {
        if flattened[i] != lookup[i - 1] {
            diff = append(diff, i)
		}
    }
    fmt.Println("diff/len", len(diff), "snip/", diff[:42])
    fmt.Println("diff/len", len(diff))

    // step - re-create an image
    bounds := pimg.Bounds()
    X /*, Y*/:= bounds.Max.X//, bounds.Max.Y
    whiteres := image.NewRGBA(bounds)
    whitedot := color.RGBA{ 255, 255, 255, 255 }
    draw.Draw(whiteres, bounds, & image.Uniform{ whitedot }, image.Point{}, draw.Src)
    whiteout, _ := os.Create("lv27w.png")
    defer whiteout.Close()

    blackres := image.NewRGBA(bounds)
    blackdot := color.RGBA{ 0, 0, 0, 255 }
    draw.Draw(blackres, bounds, & image.Uniform{ blackdot }, image.Point{}, draw.Src)
    blackout, _ := os.Create("lv27b.png")
    defer blackout.Close()
    for _, i := range diff {
        x, y := i % X, i / X
        whiteres.Set(x, y, blackdot)
        blackres.Set(x, y, whitedot)
    }
    _ = png.Encode(whiteout, whiteres)
    _ = png.Encode(blackout, blackres)

    // the output image says/ not KEY word - busy?
    DIFF := make([]uint8, len(diff))
    for i, index := range diff {
        DIFF[i] = flattened[index]
    }
    fmt.Println(yell("DIFF/snip"), string(DIFF)[:42])
    buff := bytes.NewBuffer(DIFF)
    bzreader := bzip2.NewReader(buff)
    bzdata, err := ioutil.ReadAll(bzreader)
    fmt.Println(yell("err/readall"), err)
    text := string(bzdata)
    fmt.Println(yell("text/snip"), text[:42])
    words := strings.Split(text, " ")
    fmt.Println(yell("len/"), len(words))
    set := make(map[string]bool)
    unique := []string{}
    for _, word := range words {
        if !set[word] {
            set[word] = true
            unique = append(unique, word)
        }
    }
    fmt.Println(yell("unique/len"), len(unique))
    for _, word := range unique {
        fmt.Println(yell("unique/"), word)
    }

    // step - list all keywords in p3
    cmd := exec.Command("python3", "-c", "import keyword; print(keyword.kwlist)")
    output, err := cmd.Output()
    fmt.Println(yell("err/exec"), err)
    temp := string(output)
    fmt.Print(yell("exec/out "), temp)
    keywords := strings.Split(temp[2:len(temp) - 3], "', '")
    // 👆 square bracket & single quote trimmed
    fmt.Println(yell("exec/list"), keywords)

    // step - bruteforce filter
    for _, word := range unique {
        found := false
        for _, keyword := range keywords {
            if word == keyword { found = true }
        }
        if !found { fmt.Println(yell("!keyword/"), word) }
    }
}

func palettelookup(flattened []uint8, palette [256]uint8) []uint8 {
    lookup := make([]uint8, len(flattened))
    for i, b := range flattened { lookup[i] = palette[b] }
    return lookup
}

func rgb8bit(c color.Color) (byte, byte, byte) {
    r,g,b,_ := c.RGBA()
    return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}

func init(){
    data, _ = getbody("speedboat.html")
    fmt.Println(string(data))
    yell("body ends/\n")
    re := regexp.MustCompile(`(?s)<img src="(.*?).jpg"`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(cyan("sub/"), sub)
    data, _ = getbody(sub + ".gif")
    fmt.Println("data/snippet", data[:42], "len/", len(data))
    fmt.Println("data/str", string(data)[:21], "len/", len(string(data)))
    fmt.Println("div3/", len(data) / 3, "mod3/", len(data) % 3)
    fmt.Println("div4/", len(data) / 4, "mod4/", len(data) % 4)
    // attempt png, reset reader-pointer, try reading as gif
    reader := bytes.NewReader(data)
    _, err := png.Decode(reader)
    fmt.Println(cyan("png/err"), err)
    // reset
    reader.Seek(0, 0)
    decoder, err := gif.Decode(reader)
    fmt.Println(cyan("gif/err"), err)
    _, ok := decoder.(*image.Paletted)
    fmt.Println(cyan("is paletted img?/"), ok)
    // res: plt
    pimg, _ = decoder.(*image.Paletted)
    fmt.Println(cyan("gif/typ"), reflect.TypeOf(pimg), pimg.Bounds())
}

func getbody(sub string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/hex/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) string { return Yell + s + Rest }
func cyan(s string) string { return Cyan + s + Rest }


