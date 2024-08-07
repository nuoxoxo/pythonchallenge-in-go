package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "os/exec"
    "bufio"
    "compress/bzip2"
	"bytes"
	"compress/zlib"
    "reflect"
)

const StartByte string = "bytes=1152983631-"
const URL string = "http://www.pythonchallenge.com/pc/hex/unreal.jpg"
const KEY = "redavni"
const DIR = "pc21/"
const Yell string = "\033[33m" 
const Cyan string = "\033[36m" 
const Rest string = "\033[0m"

func main(){

    // upzip w/ password
    cmd := exec.Command("unzip", "-P", KEY, DIR + "readme", "-d", DIR)
    err := cmd.Run()
    fmt.Println("err/exec", err)

    // level instruction
    f, _ := os.Open(DIR + "readme.txt")
    defer f.Close()
    scanner := bufio.NewScanner( f )
    for scanner.Scan(){
        fmt.Println("readme/", Cyan + scanner.Text() + Rest)
    }

    // open the compressed `package.pack`
    f, _ = os.Open(DIR + "package.pack")
    defer f.Close()

    PACK, _ := ioutil.ReadAll( f )
    fmt.Println("PACK/type", reflect.TypeOf( PACK ))

    pack := PACK
    res := ""
    var zlib_standard_magic_num []byte = []byte{ 0x78, 0x9c }
    var zlib_reversed_magic_num []byte = []byte{ 0x9c, 0x78 }
    var bzip2_magic_num []byte = []byte("BZh")

    for {

        if bytes.HasPrefix(pack, zlib_standard_magic_num ){
            //fmt.Println("zlib")
            reader, _ := zlib.NewReader(bytes.NewReader(pack))
            pack, _ = ioutil.ReadAll(reader)
            res += "@"
        } else if bytes.HasSuffix( pack, zlib_reversed_magic_num ){
            //fmt.Println("zlib/rev")
            for l, r := 0, len(pack) - 1; l < r; l, r = l + 1, r - 1 {
                pack[l], pack[r] = pack[r], pack[l]
            }
            res += "\n"
        } else if bytes.HasPrefix(pack, bzip2_magic_num ) {
            //fmt.Println("bzip2")
            reader := bzip2.NewReader(bytes.NewReader(pack))
            pack, err = ioutil.ReadAll(reader)
            res += " "
        } else {
            break
            panic("wtf/")
        } 
    }

    // Print the final data
    fmt.Println("\n" + Yell + res + Rest + "\n" )

}

func init(){

    // mkdir $DIR
    if _, err := os.Stat( DIR ); os.IsNotExist(err) {
        _ = os.Mkdir(DIR, os.ModePerm)
    }

    // get the file - repeat puzzle 20
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL, nil)
    req.Header.Set("Range", StartByte)
    req.SetBasicAuth("butter", "fly")
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    for k, v := range resp.Header { fmt.Println("head/", k, v) }
    _ = os.WriteFile( DIR + "readme" , body, 0644)
}



