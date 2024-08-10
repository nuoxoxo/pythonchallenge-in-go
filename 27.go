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
    raw := imgToBytes(pimg)
	trans := translateBytes(raw, palette)
    for i := 1; i < len(raw); i++ {
        if raw[i] != trans[i-1] {
            diff = append(diff, i)
		}
    }
    fmt.Println("diff/len", len(diff), "snip/", diff[:42])
    // does not work...
    /*
    diff := make(map[uint8][2]uint8)
    np := 1024 // next pixel
    for i := 0; i < len(pimg.Pix); i++ {
        index := pimg.Pix[i]
        curr := palette[int(index)]
        if np != 1024 && uint8(np) != curr {
            diff[uint8(i)] = [2]uint8{uint8(np), curr}
        }
        np = int(curr)
    }*/
    fmt.Println("diff/len", len(diff))
    bounds := pimg.Bounds()
    X/*, Y*/ := bounds.Max.X//, bounds.Max.Y
    res := image.NewRGBA(bounds)
    blackdot := color.RGBA{ 0, 0, 0, 255 }
    whitedot := color.RGBA{ 255, 255, 255, 255 }
    draw.Draw(res, bounds, & image.Uniform{ whitedot }, image.Point{}, draw.Src)
    outfile, _ := os.Create("lv27.png")
    defer outfile.Close()
    for _, i := range diff {
        x, y := i % X, i / X
        res.Set(x, y, blackdot)
    }
    _ = png.Encode(outfile, res)
}

func translateBytes(raw []byte, palette [256]uint8) []byte {
	trans := make([]byte, len(raw))
	for i, b := range raw {
		trans[i] = palette[b]
	}
	return trans
}

func imgToBytes(img image.Image) []byte {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	raw := make([]byte, 0, width*height)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			raw = append(raw, uint8(r>>8))
		}
	}

	return raw
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


