package main

import (
    "reflect"
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "os/exec"
)


const URL string = "http://www.pythonchallenge.com/pc/hex/"
const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"


func main(){
    res, err := exec.Command("python3", "-c", "import this").Output()
    fmt.Println("err/exec", err)
    fmt.Println("typ/", reflect.TypeOf(res))
    fmt.Println("\n" + Yell + string(res) + Rest)

    source := string(res)
    re := regexp.MustCompile(`(?is)in the face of (.*?)[, ]`)
    matches := re.FindAllStringSubmatch( source, -1)
    word := matches[0][1]
    fmt.Println("word/", Cyan + word, Rest)
}

func init(){

    // GET
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + "bonus.html", nil)
    req.SetBasicAuth( "butter", "fly" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    temp, _ := ioutil.ReadAll(resp.Body)
    body := string(temp)

    fmt.Println(Cyan + body + Rest + "body ends/\n")

    re := regexp.MustCompile(`(?s)<!--\n'(.*?)'\n-->`)
    matches := re.FindAllStringSubmatch(body, -1)
    msg := matches[0][1]
    fmt.Println("cmt/", Yell + msg, Rest)
    fmt.Println("r13/", Cyan + r13(msg), Rest)
}

func r13(s string)string{
    res := ""
    for _, char := range s {
        if char < 'a' || char > 'z' {
            res += string(char)
        } else {
            res += string((char - 'a' + 13) % 26 + 'a')
        }
    }
    return res
}


