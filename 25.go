package main

import (
    "fmt"
    "net/http"
    _"io"
    "io/ioutil"
    "strconv"
    "reflect"
)

func main(){
    N := getlastindex(1, ".wav")
    waves := getwaves(N, ".wav")
    
    // one long continous data w/o wave header
    res := []uint8{}
    idx := 44
    for _, wave := range waves {
        data := wave[idx:]
        res = append(res, data...)
    }
    fmt.Println(len(res))
}

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
        data, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("len/typ", len(data), reflect.TypeOf(data), i)
    }
    return i - 1
}

const URL string = "http://www.pythonchallenge.com/pc/hex/lake"//1.jpg"
const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) { fmt.Println(Yell + s + Rest) }
func cy(s string)   { fmt.Println(Cyan + s + Rest) }


